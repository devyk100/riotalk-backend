package state

import (
	"chat-server/types"
	"github.com/gorilla/websocket"
)

var Clients = make(map[int64]*websocket.Conn)
var IsChannelActive = make(map[int64]bool)
var IsUserActive = make(map[int64]bool)
var Events = make(map[int64]chan *types.Event)
