package cmd

import (
	"chat-server/state"
	"chat-server/types"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

func MainReadLoop(conn *websocket.Conn, userId int64) {
	defer func() {
		conn.Close()
		close(state.Events[userId])
	}()
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Client disconnected")
			break
		}
		var ReceivedMesssage types.Event
		err = json.Unmarshal(msg, &ReceivedMesssage)
		if err != nil {
			fmt.Println("Error unmarshalling message:", err)
			continue
		}

		switch ReceivedMesssage.Event {
		case "chat":
			// appropriately handle channel, and user mesasages differently
			// publish them to the pubsub respective channel
			// If the pubsub channel is not active, make it active
		case "auth":
		case "close":
			// exit and close the connection
		default:
		}
	}
}
