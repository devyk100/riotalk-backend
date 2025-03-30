package cmd

import (
	"chat-server/state"
	"chat-server/types"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

func MainWriteLoop(conn *websocket.Conn, userId int64) {
	for event := range state.Events[userId] {

		response := types.Event{
			Event: "",
			Type:  "",
			To:    0,
		}

		responseJson, err := json.Marshal(response)
		if err != nil {
			fmt.Println("Error marshalling response:", err)
		}
		err = conn.WriteMessage(websocket.TextMessage, responseJson)
	}
}
