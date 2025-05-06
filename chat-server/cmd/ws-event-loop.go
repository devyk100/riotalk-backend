package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"shared/redis"
	"shared/types"
)

func TopicEventHandlerCallback(Client *types.Client, val string, userId int64) error {

	var event types.Event
	err := json.Unmarshal([]byte(val), &event)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	switch event.Event {
	case "chat":
		if event.FromID == userId {
			return errors.New("user already in chat")
		}
		err := Client.SafeWriteMessage(websocket.TextMessage, []byte(val))
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	default:
		fmt.Println(event.Event, "IS AN INVALID EVENT")
	}
	return nil
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
		err := TopicEventHandlerCallback(client, msg.Payload, userId)
		// This is important, when the websocket closes, or so, you must
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	fmt.Println("Exiting the ws event loop")
}
