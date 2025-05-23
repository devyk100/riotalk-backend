package cmd

import (
	"chat-server/state"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"shared/types"
	"shared/utils"
)

func HandleWSAuth(conn *websocket.Conn) (int64, error) {
	_, msg, err := conn.ReadMessage()
	if err != nil {
		fmt.Println("Client disconnected")
		return -1, err
	}
	var ReceivedAuthMessage types.Event
	err = json.Unmarshal(msg, &ReceivedAuthMessage)
	if err != nil {
		fmt.Println("Error unmarshalling message:", err)
		return -1, err
	}
	if ReceivedAuthMessage.Event == "auth" {
		fmt.Println("Received auth message")
	} else {
		return 0, fmt.Errorf("invalid event")
	}
	_, method, userId, err := utils.ParseToken(ReceivedAuthMessage.Token)
	state.AccessTokens[userId] = ReceivedAuthMessage.Token
	if err != nil {
		return 0, err
	}
	fmt.Println("Was authorized using", method, "user", userId)
	return userId, nil
}
