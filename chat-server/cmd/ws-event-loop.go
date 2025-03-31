package cmd

import (
	"chat-server/redis"
	"chat-server/state"
	"chat-server/types"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

func TopicEventHandlerCallback(val string, userId int64) {
	var event types.Event
	err := json.Unmarshal([]byte(val), &event)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	switch event.Event {
	case "chat":
		if state.Clients[userId] == nil {
			return
		}
		if event.From == userId {
			return
		}
		err := state.Clients[userId].SafeWriteMessage(websocket.TextMessage, []byte(val))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	default:
		fmt.Println(event.Event, "IS AN INVALID EVENT")
	}
}

func MainEventLoop(client *types.Client, userId int64) {
	serverList, err := FetchServerList(userId)
	if err != nil {
		fmt.Println(err.Error())
		client.Conn.Close()
		return
	}
	fmt.Println(serverList, "Is teh list of servers connected to")

	// A NON DUPLICATE SERVER LIST
	topicKeySet := make(map[string]struct{})

	for _, server := range serverList {
		topicKeySet[redis.ServerKey(server.ID)] = struct{}{}
	}

	topicKeySet[redis.UserKey(userId)] = struct{}{}

	// Converting this map keys back to a slice, to iterate simply
	topicKeyList := make([]string, 0, len(topicKeySet))
	for key := range topicKeySet {
		topicKeyList = append(topicKeyList, key)
	}

	sub := redis.RedisClient.Subscribe(context.Background(), topicKeyList...)
	ch := sub.Channel()
	for msg := range ch {
		TopicEventHandlerCallback(msg.Payload, userId)
	}
}
