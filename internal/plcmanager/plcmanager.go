package plcmanager

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
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

// runTag executa a coleta contínua de uma tag. Se ocorrer um erro crítico, envia o erro pelo canal errChan.
func runTag(ctx context.Context, plcID int, tag database.Tag, client *plc.Client, redis *cache.RedisCache, config TagConfig, logger *database.Logger, errChan chan error) {
	ticker := time.NewTicker(config.ScanRate)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Printf("Encerrando monitoramento da tag %s (PLC %d)", tag.Name, plcID)
			logger.Warn("Encerramento do monitoramento da tag", tag.Name)
			return
		case <-ticker.C:
			value, err := client.ReadTag(tag.DBNumber, tag.ByteOffset, tag.DataType)
			if err != nil {
				log.Printf("Erro ao ler tag %s no PLC %d: %v", tag.Name, plcID, err)
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
				log.Printf("Erro ao atualizar Redis para tag %s no PLC %d: %v", tag.Name, plcID, err)
				logger.Error("Erro ao atualizar Redis para tag", fmt.Sprintf("%s: %v", tag.Name, err))
				if isCriticalError(err) {
					errChan <- fmt.Errorf("erro crítico ao atualizar Redis para tag %s: %v", tag.Name, err)
					return
				}
			} else {
				log.Printf("PLC %d - Tag %s atualizada: %v", plcID, tag.Name, value)
			}
		}
	}
}

// managePLCTags gerencia as goroutines de coleta de tags de um PLC.
func managePLCTags(ctx context.Context, plcID int, db *database.DB, client *plc.Client, redis *cache.RedisCache, logger *database.Logger) error {
	tagRunners := make(map[int]TagRunner)
	reloadTicker := time.NewTicker(5 * time.Second)
	defer reloadTicker.Stop()

	errChan := make(chan error, 10)

	for {
		select {
		case <-ctx.Done():
			log.Printf("Shutdown detectado, encerrando gerenciamento de tags do PLC %d", plcID)
			logger.Warn("Shutdown do gerenciamento de tags", fmt.Sprintf("PLC ID: %d", plcID))
			return nil
		case err := <-errChan:
			return err
		default:
		}

		tags, err := db.GetPLCTags(plcID)
		if err != nil {
			log.Printf("Erro ao carregar tags do PLC %d: %v", plcID, err)
			logger.Error("Erro ao carregar tags", fmt.Sprintf("PLC ID %d: %v", plcID, err))
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
				go runTag(childCtx, plcID, tag, client, redis, newConfig, logger, errChan)
				log.Printf("Iniciou monitoramento da tag %s no PLC %d", tag.Name, plcID)
				logger.Info("Iniciou monitoramento da tag", tag.Name)
			} else if runner.config != newConfig {
				runner.cancel()
				childCtx, cancel := context.WithCancel(ctx)
				tagRunners[tagID] = TagRunner{cancel: cancel, config: newConfig}
				go runTag(childCtx, plcID, tag, client, redis, newConfig, logger, errChan)
				log.Printf("Reiniciou monitoramento da tag %s no PLC %d", tag.Name, plcID)
				logger.Info("Reiniciou monitoramento da tag", tag.Name)
			}
		}

		for tagID, runner := range tagRunners {
			if _, exists := activeTags[tagID]; !exists {
				runner.cancel()
				delete(tagRunners, tagID)
				log.Printf("Encerrado monitoramento da tag %d no PLC %d", tagID, plcID)
				logger.Warn("Encerrado monitoramento da tag", fmt.Sprintf("ID: %d", tagID))
			}
		}
		<-reloadTicker.C
	}
}

// updatePLCStatus verifica periodicamente a conectividade do PLC usando client.Ping()
// e atualiza os campos status e last_update.
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

			// Tenta atualizar a tabela; mesmo se rowsAffected for 0, a função UpdatePLCStatus retorna nil.
			if err := db.UpdatePLCStatus(status); err != nil {
				logger.Error("Erro ao atualizar status do PLC", fmt.Sprintf("PLC ID %d: %v", plcID, err))
			} else {
				if status.Status != lastStatus {
					logger.Info("Status do PLC atualizado", fmt.Sprintf("PLC ID %d: %s", plcID, status.Status))
					lastStatus = status.Status
				}
			}

			if failCount >= threshold {
				logger.Warn("Muitas falhas consecutivas de ping", fmt.Sprintf("PLC %d: %d falhas consecutivas", plcID, failCount))
				// Aqui podemos reiniciar o contador para continuar tentando atualizar.
				failCount = 0
			}
		}
	}
}

// runPLC tenta manter a conexão com um PLC.
func runPLC(ctx context.Context, plcConfig database.PLC, db *database.DB, redis *cache.RedisCache, logger *database.Logger) {
	for {
		select {
		case <-ctx.Done():
			log.Printf("Shutdown detectado para PLC %s. Encerrando rotina.", plcConfig.Name)
			return
		default:
		}

		client, err := plc.NewClient(plcConfig.IPAddress, plcConfig.Rack, plcConfig.Slot)
		if err != nil {
			logger.Error("Erro ao conectar ao PLC", fmt.Sprintf("%s: %v", plcConfig.Name, err))
			// Atualiza o status para offline quando não consegue conectar
			offlineStatus := database.PLCStatus{
				PLCID:      plcConfig.ID,
				Status:     "offline",
				LastUpdate: time.Now(),
			}
			if errUpd := db.UpdatePLCStatus(offlineStatus); errUpd != nil {
				logger.Error("Erro ao atualizar status para offline", fmt.Sprintf("PLC ID %d: %v", plcConfig.ID, errUpd))
			}
			time.Sleep(10 * time.Second)
			continue
		}
		log.Printf("Conectado ao PLC: %s (%s)", plcConfig.Name, plcConfig.IPAddress)
		logger.Info("Conectado ao PLC", plcConfig.Name)

		plcCtx, cancel := context.WithCancel(ctx)
		var wg sync.WaitGroup
		errChan := make(chan error, 2)
		wg.Add(2)
		go func() {
			defer wg.Done()
			if err := updatePLCStatus(plcCtx, plcConfig.ID, client, db, logger); err != nil {
				errChan <- err
			}
		}()
		go func() {
			defer wg.Done()
			if err := managePLCTags(plcCtx, plcConfig.ID, db, client, redis, logger); err != nil {
				errChan <- err
			}
		}()

		select {
		case err := <-errChan:
			logger.Error("Erro crítico no PLC", fmt.Sprintf("%s: %v", plcConfig.Name, err))
			// Força atualização para offline
			offlineStatus := database.PLCStatus{
				PLCID:      plcConfig.ID,
				Status:     "offline",
				LastUpdate: time.Now(),
			}
			if errUpd := db.UpdatePLCStatus(offlineStatus); errUpd != nil {
				logger.Error("Erro ao atualizar status para offline", fmt.Sprintf("PLC ID %d: %v", plcConfig.ID, errUpd))
			}
		case <-plcCtx.Done():
			// Opcional: atualizar para offline aqui também
		}
		cancel()
		wg.Wait()
		client.Close()
		log.Printf("Conexão com PLC %s perdida. Tentando reconectar...", plcConfig.Name)
		logger.Warn("Conexão perdida", plcConfig.Name)
		time.Sleep(5 * time.Second)
	}
}

// RunAllPLCs consulta os PLCs ativos e inicia uma rotina de gerenciamento para cada um.
func RunAllPLCs(ctx context.Context, db *database.DB, redis *cache.RedisCache, logger *database.Logger) {
	plcs, err := db.GetActivePLCs()
	if err != nil {
		log.Fatalf("Erro ao carregar PLCs: %v", err)
	}
	var wg sync.WaitGroup
	for _, plcConfig := range plcs {
		wg.Add(1)
		go func(p database.PLC) {
			defer wg.Done()
			runPLC(ctx, p, db, redis, logger)
		}(plcConfig)
	}
	wg.Wait()
}
