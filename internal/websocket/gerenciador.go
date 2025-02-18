package websocket

import (
	"log"
	"sync"
	"time"

	"plc_project/internal/cache"
	"plc_project/internal/database"

	"github.com/gorilla/websocket"
)

type Gerenciador struct {
	clientes     map[*Cliente]bool
	broadcast    chan MensagemWS
	registrar    chan *Cliente
	desregistrar chan *Cliente
	db           *database.DB
	redis        *cache.RedisCache
	mutex        sync.RWMutex
}

type Cliente struct {
	conn        *websocket.Conn
	gerenciador *Gerenciador
	enviar      chan MensagemWS
}

func NovoGerenciador(db *database.DB, redis *cache.RedisCache) *Gerenciador {
	return &Gerenciador{
		clientes:     make(map[*Cliente]bool),
		broadcast:    make(chan MensagemWS),
		registrar:    make(chan *Cliente),
		desregistrar: make(chan *Cliente),
		db:           db,
		redis:        redis,
	}
}

func (g *Gerenciador) Iniciar() {
	go g.coletarDados()

	for {
		select {
		case cliente := <-g.registrar:
			g.mutex.Lock()
			g.clientes[cliente] = true
			log.Printf("Novo cliente conectado. Total: %d", len(g.clientes))
			g.mutex.Unlock()

		case cliente := <-g.desregistrar:
			g.mutex.Lock()
			if _, ok := g.clientes[cliente]; ok {
				delete(g.clientes, cliente)
				close(cliente.enviar)
				log.Printf("Cliente desconectado. Total: %d", len(g.clientes))
			}
			g.mutex.Unlock()

		case mensagem := <-g.broadcast:
			g.mutex.RLock()
			for cliente := range g.clientes {
				select {
				case cliente.enviar <- mensagem:
				default:
					close(cliente.enviar)
					delete(g.clientes, cliente)
				}
			}
			g.mutex.RUnlock()
		}
	}
}

func (g *Gerenciador) coletarDados() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		plcs, err := g.db.GetActivePLCs()
		if err != nil {
			log.Printf("Erro ao buscar PLCs: %v", err)
			continue
		}

		for _, plc := range plcs {
			tags, err := g.db.GetPLCTags(plc.ID)
			if err != nil {
				log.Printf("Erro ao buscar tags do PLC %s (ID=%d): %v", plc.Name, plc.ID, err)
				continue
			}

			msg := MensagemWS{
				PLC: StatusPLC{
					ID:                plc.ID,
					Name:              plc.Name,
					Status:            plc.Status,
					UltimaAtualizacao: plc.LastUpdate.Format(time.RFC3339),
				},
				Tags: make([]ValorTag, 0, len(tags)),
			}

			for _, tag := range tags {
				valor, err := g.redis.GetTagValue(plc.ID, tag.ID)
				if err != nil {
					msg.Tags = append(msg.Tags, ValorTag{
						ID:    tag.ID,
						Nome:  tag.Name,
						Valor: nil,
					})
					continue
				}

				msg.Tags = append(msg.Tags, ValorTag{
					ID:    tag.ID,
					Nome:  tag.Name,
					Valor: valor.Value,
				})
			}

			if len(msg.Tags) > 0 {
				g.mutex.RLock()
				for cliente := range g.clientes {
					select {
					case cliente.enviar <- msg:
					default:
						close(cliente.enviar)
						delete(g.clientes, cliente)
					}
				}
				g.mutex.RUnlock()
			}
		}
	}
}
