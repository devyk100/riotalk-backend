package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"shared/redis"
	"shared/sqs"
	"shared/types"
)

func ClientEventHandler(client *types.Client, ctx context.Context, msg []byte, userId int64) {
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
			redis.PushToRecentMessages(ctx, redis.RecentMessageServerKey(ReceivedMessage.To, ReceivedMessage.ChannelId), string(msg))
			err = sqs.SendMessage(string(msg), sqs.MESSAGE_GROUP_SERVER)
			if err != nil {
				fmt.Println("Error sending SQS message:", err)
				return
			}
		case "user":
			err := redis.Publish(ctx, redis.UserKey(ReceivedMessage.To), string(msg))
			if err != nil {
				fmt.Println("Error publishing:", err)
				return
			}
			redis.PushToRecentMessages(ctx, redis.RecentMessageUserKey(ReceivedMessage.To, userId), string(msg))
			err = sqs.SendMessage(string(msg), sqs.MESSAGE_GROUP_USER)
			if err != nil {
				fmt.Println("Error sending SQS message:", err)
				return
			}
		default:
			fmt.Println(ReceivedMessage.Event, "IS AN INVALID EVENT")
		}
	case "history":
		var Key string
		switch ReceivedMessage.Type {
		case "server":
			Key = redis.RecentMessageServerKey(ReceivedMessage.Of, ReceivedMessage.ChannelId)
		case "user":
			Key = redis.RecentMessageUserKey(ReceivedMessage.Of, userId)
		default:
			fmt.Println("Error type:", ReceivedMessage.Type)
			return
		}
		messages, err := redis.GetRecentMessages(ctx, Key)
		if err != nil {
			return
		}
		var events []types.Event
		for _, raw := range messages {
			var event types.Event
			err := json.Unmarshal([]byte(raw), &event)
			if err != nil {
				fmt.Println("Error unmarshalling message:", err)
				return
			}
			events = append(events, event)
		}
		finalJSON, err := json.Marshal(events)
		if err != nil {
			fmt.Println("Error marshalling message:", err)
			return
		}
		err = client.SafeWriteMessage(websocket.TextMessage, finalJSON)
		if err != nil {
			fmt.Println("Error writing message:", err)
			return
		}
		return

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
		ClientEventHandler(client, ctx, msg, userId)
	}
}
