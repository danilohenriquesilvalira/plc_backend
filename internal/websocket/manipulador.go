package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (g *Gerenciador) ManipularWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Erro ao abrir conexão WebSocket: %v", err)
		http.Error(w, "Erro ao abrir conexão WebSocket", http.StatusBadRequest)
		return
	}

	cliente := &Cliente{
		conn:        conn,
		gerenciador: g,
		enviar:      make(chan MensagemWS),
	}

	g.registrar <- cliente

	go cliente.bombearEscrita()
	go cliente.bombearLeitura()
}

func (c *Cliente) bombearEscrita() {
	defer func() {
		c.conn.Close()
	}()

	for mensagem := range c.enviar {
		if err := c.conn.WriteJSON(mensagem); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Erro ao enviar mensagem: %v", err)
			}
			return
		}
	}
}

func (c *Cliente) bombearLeitura() {
	defer func() {
		c.gerenciador.desregistrar <- c
		c.conn.Close()
	}()

	for {
		_, _, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Erro na conexão: %v", err)
			}
			break
		}
	}
}
