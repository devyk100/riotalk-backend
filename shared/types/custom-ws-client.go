package types

import (
	"github.com/gorilla/websocket"
	"sync"
)

type Client struct {
	Conn *websocket.Conn
	Mu   *sync.Mutex // Protects writes
}

func (c Client) SafeWriteMessage(messageType int, message []byte) error {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	return c.Conn.WriteMessage(messageType, message)
}
