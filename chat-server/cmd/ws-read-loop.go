package cmd

import (
	"chat-server/redis"
	"chat-server/types"
	"context"
	"encoding/json"
	"fmt"
)

func ClientEventHandler(ctx context.Context, msg []byte, userId int64) {
	var ReceivedMessage types.Event
	err := json.Unmarshal(msg, &ReceivedMessage)
	if err != nil {
		fmt.Println("Error unmarshalling message:", err)
		return
	}
	switch ReceivedMessage.Event {
	case "chat":
		switch ReceivedMessage.Type {
		case "server":
			err := redis.Publish(ctx, redis.ServerKey(ReceivedMessage.To), string(msg))
			if err != nil {
				fmt.Println("Error publishing:", err)
				return
			}
		case "user":
			err := redis.Publish(ctx, redis.UserKey(ReceivedMessage.To), string(msg))
			if err != nil {
				fmt.Println("Error publishing:", err)
				return
			}
		default:
			fmt.Println(ReceivedMessage.Event, "IS AN INVALID EVENT")
		}
	case "close":
	default:
		fmt.Println(ReceivedMessage.Event, "IS AN INVALID EVENT")
	}
}

func MainReadLoop(client *types.Client, userId int64) {
	ctx := context.Background()
	defer func() {
		client.Conn.Close()
	}()
	for {
		_, msg, err := client.Conn.ReadMessage() // Blocking
		if err != nil {
			fmt.Println("Client disconnected")
			break
		}
		ClientEventHandler(ctx, msg, userId)
	}
}
