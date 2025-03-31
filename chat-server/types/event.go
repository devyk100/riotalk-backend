package types

import (
	"github.com/gorilla/websocket"
	"sync"
)

type Event struct {
	Event     string `json:"event"`      // chat, auth, active
	Type      string `json:"type"`       // server, user
	To        int64  `json:"to"`         // server_id, user_id
	ChannelId int64  `json:"channel_id"` // channel_id, if To is a server_id
	Token     string `json:"token"`      // in auth
	From      int64  `json:"from"`       // from user_id, at the user end or in a channel
	Payload   string `json:"payload"`
}
type Client struct {
	Conn *websocket.Conn
	Mu   sync.Mutex // Protects writes
}

func (c *Client) SafeWriteMessage(messageType int, message []byte) error {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	return c.Conn.WriteMessage(messageType, message)
}
