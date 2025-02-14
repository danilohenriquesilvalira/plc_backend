package websocket

// MensagemWS representa a mensagem completa enviada pelo WebSocket
type MensagemWS struct {
	PLC  StatusPLC  `json:"plc"`
	Tags []ValorTag `json:"tags"`
}

// StatusPLC representa o status do PLC
type StatusPLC struct {
	ID                int    `json:"plc_id"`
	Status            string `json:"status"`
	UltimaAtualizacao string `json:"last_update"`
}

// ValorTag representa o valor de uma tag
type ValorTag struct {
	ID    int         `json:"id"`
	Nome  string      `json:"name"`
	Valor interface{} `json:"value"`
}
