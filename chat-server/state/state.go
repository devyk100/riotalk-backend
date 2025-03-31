package state

import (
	"chat-server/types"
)

var Clients = make(map[int64]*types.Client)
var AccessTokens = make(map[int64]string)
