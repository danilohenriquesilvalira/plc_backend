package plcmanager

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"plc_project/internal/cache"
	"plc_project/internal/database"
	"plc_project/internal/plc"
)

// TagConfig guarda os parâmetros para o monitoramento de uma tag.
type TagConfig struct {
	ScanRate       time.Duration
	MonitorChanges bool
}

// TagRunner guarda o cancelamento da goroutine de coleta e a configuração aplicada.
type TagRunner struct {
	cancel context.CancelFunc
	config TagConfig
}

// isCriticalError verifica se o erro indica perda de conexão (crítico).
func isCriticalError(err error) bool {
	if err == nil {
		return false
	}
	lower := strings.ToLower(err.Error())
	return strings.Contains(lower, "forçado") || strings.Contains(lower, "cancelado")
}

// runTag executa a coleta contínua de uma tag.
func runTag(ctx context.Context, plcID int, plcName string, tag database.Tag, client *plc.Client, redis *cache.RedisCache, config TagConfig, logger *database.Logger, errChan chan error) {
	ticker := time.NewTicker(config.ScanRate)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Printf("Encerrando monitoramento da tag %s (%s)", tag.Name, plcName)
			logger.Warn("Encerramento do monitoramento da tag", tag.Name)
			return
		case <-ticker.C:
			value, err := client.ReadTag(tag.DBNumber, tag.ByteOffset, tag.DataType)
			if err != nil {
				log.Printf("Erro ao ler tag %s no PLC %s: %v", tag.Name, plcName, err)
				logger.Error("Erro ao ler tag", fmt.Sprintf("%s: %v", tag.Name, err))
				if isCriticalError(err) {
					errChan <- fmt.Errorf("erro crítico na tag %s: %v", tag.Name, err)
					return
				}
				continue
			}
			if config.MonitorChanges {
				oldValue, err := redis.GetTagValue(plcID, tag.ID)
				if err == nil && plc.CompareValues(oldValue.Value, value) {
					continue
				}
			}
			if err := redis.SetTagValue(plcID, tag.ID, value); err != nil {
				log.Printf("Erro ao atualizar Redis para tag %s no PLC %s: %v", tag.Name, plcName, err)
				logger.Error("Erro ao atualizar Redis para tag", fmt.Sprintf("%s: %v", tag.Name, err))
				if isCriticalError(err) {
					errChan <- fmt.Errorf("erro crítico ao atualizar Redis para tag %s: %v", tag.Name, err)
					return
				}
			} else {
				log.Printf("%s - Tag %s atualizada: %v", plcName, tag.Name, value)
			}
		}
	}
}

// managePLCTags gerencia as goroutines de coleta de tags de um PLC.
func managePLCTags(ctx context.Context, plcID int, plcName string, db *database.DB, client *plc.Client, redis *cache.RedisCache, logger *database.Logger) error {
	tagRunners := make(map[int]TagRunner)
	reloadTicker := time.NewTicker(5 * time.Second)
	defer reloadTicker.Stop()

	errChan := make(chan error, 10)

	for {
		select {
		case <-ctx.Done():
			log.Printf("Shutdown detectado, encerrando gerenciamento de tags do PLC %s", plcName)
			logger.Warn("Shutdown do gerenciamento de tags", fmt.Sprintf("PLC: %s", plcName))
			return nil
		case err := <-errChan:
			return err
		default:
		}

		tags, err := db.GetPLCTags(plcID)
		if err != nil {
			log.Printf("Erro ao carregar tags do PLC %s: %v", plcName, err)
			logger.Error("Erro ao carregar tags", fmt.Sprintf("PLC %s: %v", plcName, err))
			time.Sleep(5 * time.Second)
			continue
		}

		activeTags := make(map[int]database.Tag)
		for _, tag := range tags {
			if tag.Active {
				activeTags[tag.ID] = tag
			}
		}

		for tagID, tag := range activeTags {
			newConfig := TagConfig{
				ScanRate:       time.Duration(tag.ScanRate) * time.Millisecond,
				MonitorChanges: tag.MonitorChanges,
			}
			runner, exists := tagRunners[tagID]
			if !exists {
				childCtx, cancel := context.WithCancel(ctx)
				tagRunners[tagID] = TagRunner{cancel: cancel, config: newConfig}
				go runTag(childCtx, plcID, plcName, tag, client, redis, newConfig, logger, errChan)
				log.Printf("Iniciou monitoramento da tag %s no PLC %s", tag.Name, plcName)
				logger.Info("Iniciou monitoramento da tag", tag.Name)
			} else if runner.config != newConfig {
				runner.cancel()
				childCtx, cancel := context.WithCancel(ctx)
				tagRunners[tagID] = TagRunner{cancel: cancel, config: newConfig}
				go runTag(childCtx, plcID, plcName, tag, client, redis, newConfig, logger, errChan)
				log.Printf("Reiniciou monitoramento da tag %s no PLC %s", tag.Name, plcName)
				logger.Info("Reiniciou monitoramento da tag", tag.Name)
			}
		}

		for tagID, runner := range tagRunners {
			if _, exists := activeTags[tagID]; !exists {
				runner.cancel()
				delete(tagRunners, tagID)
				log.Printf("Encerrado monitoramento da tag %d no PLC %s", tagID, plcName)
				logger.Warn("Encerrado monitoramento da tag", fmt.Sprintf("ID: %d", tagID))
			}
		}
		<-reloadTicker.C
	}
}

// updatePLCStatus verifica periodicamente a conectividade do PLC.
func updatePLCStatus(ctx context.Context, plcID int, client *plc.Client, db *database.DB, logger *database.Logger) error {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	failCount := 0
	threshold := 3
	var lastStatus string

	for {
		select {
		case <-ctx.Done():
			log.Printf("Encerrando atualização de status do PLC %d", plcID)
			return nil
		case <-ticker.C:
			status := database.PLCStatus{
				PLCID:      plcID,
				LastUpdate: time.Now(),
			}
			if err := client.Ping(); err != nil {
				status.Status = "offline"
				failCount++
				logger.Error("Ping falhou para PLC", fmt.Sprintf("PLC ID %d: %v", plcID, err))
			} else {
				status.Status = "online"
				failCount = 0
			}

			if err := db.UpdatePLCStatus(status); err != nil {
				logger.Error("Erro ao atualizar status do PLC", fmt.Sprintf("PLC ID %d: %v", plcID, err))
			} else if status.Status != lastStatus {
				logger.Info("Status do PLC atualizado", fmt.Sprintf("PLC ID %d: %s", plcID, status.Status))
				lastStatus = status.Status
			}

			if failCount >= threshold {
				logger.Warn("Muitas falhas consecutivas de ping", fmt.Sprintf("PLC %d: %d falhas consecutivas", plcID, failCount))
				failCount = 0
			}
		}
	}
}

// runPLC tenta manter a conexão com um PLC.
// runPLC tenta manter a conexão com um PLC.
func runPLC(ctx context.Context, plcConfig database.PLC, db *database.DB, redis *cache.RedisCache, logger *database.Logger) {
	retryTicker := time.NewTicker(5 * time.Second)
	defer retryTicker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Printf("Encerrando monitoramento do PLC %s", plcConfig.Name)
			return
		default:
			// Verifica se o PLC ainda está ativo antes de tentar conectar
			plcs, err := db.GetActivePLCs()
			if err == nil {
				plcAtivo := false
				for _, p := range plcs {
					if p.ID == plcConfig.ID {
						plcAtivo = true
						break
					}
				}
				if !plcAtivo {
					return // Encerra a goroutine se o PLC não estiver mais ativo
				}
			}

			client, err := plc.NewClient(plcConfig.IPAddress, plcConfig.Rack, plcConfig.Slot)
			if err != nil {
				select {
				case <-ctx.Done():
					return
				default:
					logger.Error("Erro ao conectar ao PLC", fmt.Sprintf("%s: %v", plcConfig.Name, err))
					offlineStatus := database.PLCStatus{
						PLCID:      plcConfig.ID,
						Status:     "offline",
						LastUpdate: time.Now(),
					}
					if errUpd := db.UpdatePLCStatus(offlineStatus); errUpd != nil {
						logger.Error("Erro ao atualizar status offline", fmt.Sprintf("PLC ID %d: %v", plcConfig.ID, errUpd))
					}
				}

				select {
				case <-retryTicker.C:
					continue
				case <-ctx.Done():
					return
				}
			}

			log.Printf("Conectado ao PLC: %s (%s)", plcConfig.Name, plcConfig.IPAddress)
			logger.Info("Conectado ao PLC", plcConfig.Name)

			clientCtx, clientCancel := context.WithCancel(ctx)
			errChan := make(chan error, 2)

			go func() {
				if err := updatePLCStatus(clientCtx, plcConfig.ID, client, db, logger); err != nil {
					errChan <- err
				}
			}()

			go func() {
				if err := managePLCTags(clientCtx, plcConfig.ID, plcConfig.Name, db, client, redis, logger); err != nil {
					errChan <- fmt.Errorf("erro no gerenciamento de tags: %v", err)
				}
			}()

			select {
			case err := <-errChan:
				clientCancel()
				logger.Error("Erro crítico no PLC", fmt.Sprintf("%s: %v", plcConfig.Name, err))
				client.Close()
			case <-clientCtx.Done():
				clientCancel()
				client.Close()
				return
			}

			select {
			case <-retryTicker.C:
			case <-ctx.Done():
				return
			}
		}
	}
}

// RunAllPLCs consulta os PLCs ativos e inicia uma rotina de gerenciamento para cada um.
func RunAllPLCs(ctx context.Context, db *database.DB, redis *cache.RedisCache, logger *database.Logger) {
	plcCancels := make(map[int]struct {
		cancel context.CancelFunc
		config database.PLC
	})

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			for _, p := range plcCancels {
				p.cancel()
			}
			return

		case <-ticker.C:
			plcs, err := db.GetActivePLCs()
			if err != nil {
				logger.Error("Erro ao carregar PLCs", err.Error())
				continue
			}

			// Remove PLCs inativos primeiro
			for plcID, p := range plcCancels {
				found := false
				for _, plc := range plcs {
					if plc.ID == plcID {
						found = true
						break
					}
				}
				if !found {
					p.cancel()
					delete(plcCancels, plcID)
					log.Printf("Removendo monitoramento do PLC ID %d - não está mais ativo", plcID)
					logger.Info("PLC removido", fmt.Sprintf("PLC ID: %d - inativo", plcID))
				}
			}

			// Depois verifica novos PLCs ou mudanças
			for _, plc := range plcs {
				current, exists := plcCancels[plc.ID]

				if !exists || current.config.IPAddress != plc.IPAddress ||
					current.config.Rack != plc.Rack ||
					current.config.Slot != plc.Slot {

					if exists {
						current.cancel()
						delete(plcCancels, plc.ID)
						log.Printf("Reiniciando PLC %s devido a mudança de configuração", plc.Name)
						logger.Info("Reiniciando PLC", fmt.Sprintf("PLC: %s, Motivo: mudança de configuração", plc.Name))
					}

					plcCtx, cancel := context.WithCancel(ctx)
					plcCancels[plc.ID] = struct {
						cancel context.CancelFunc
						config database.PLC
					}{
						cancel: cancel,
						config: plc,
					}

					go func(p database.PLC) {
						runPLC(plcCtx, p, db, redis, logger)
					}(plc)
				}
			}
		}
	}
}
