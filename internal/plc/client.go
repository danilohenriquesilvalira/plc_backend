package plc

import (
	"fmt"
	"net"
	"time"

	"github.com/robinson/gos7"
)

type Client struct {
	client  gos7.Client
	handler *gos7.TCPClientHandler
}

func NewClient(ip string, rack, slot int) (*Client, error) {
	handler := gos7.NewTCPClientHandler(ip, rack, slot)
	handler.Timeout = 10 * time.Second

	if err := handler.Connect(); err != nil {
		return nil, err
	}

	return &Client{
		client:  gos7.NewClient(handler),
		handler: handler,
	}, nil
}

func (c *Client) Close() {
	if c.handler != nil {
		c.handler.Close()
	}
}

// Ping tenta estabelecer uma conexão TCP para testar a conectividade.
// Se c.handler.Address não tiver a porta, adiciona ":102".
func (c *Client) Ping() error {
	address := c.handler.Address
	if _, _, err := net.SplitHostPort(address); err != nil {
		address = fmt.Sprintf("%s:102", address)
	}
	conn, err := net.DialTimeout("tcp", address, 3*time.Second)
	if err != nil {
		return fmt.Errorf("falha no ping TCP: %v", err)
	}
	conn.Close()
	return nil
}
