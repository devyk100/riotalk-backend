package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/joho/godotenv"
	"log"
	"persist-worker/db"
	chatserversqs "persist-worker/sqs"
	"persist-worker/types"
	"time"
)

var PERSIST_FREQUENCY = 1 * time.Second

// IN PRODUCTION, TRY PERSISTING AGAIN, EVEN WHEN FAILED
func PersistWorker() {
	ctx := context.Background()
	err := db.InitDb(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	UserQueries := db.BatchInsertUserToUserChatParams{
		Content:    make([]string, 0),
		FromUserID: make([]int64, 0),
		ToUserID:   make([]int64, 0),
		Type:       make([]string, 0),
		TimeAt:     make([]int64, 0),
	}
	ServerQueries := db.BatchInsertUserToChannelChatParams{
		Content:    make([]string, 0),
		ReplyOf:    make([]int64, 0),
		FromUserID: make([]int64, 0),
		ChannelID:  make([]int64, 0),
		Type:       make([]string, 0),
		TimeAt:     make([]int64, 0),
	}
	for {
		messages, err := chatserversqs.ReceiveMessage()
		if err != nil {
			log.Printf("Error receiving messages: %v\n", err)
			return
		}
		for _, message := range messages {
			fmt.Println(*message.Body, "is one of the messages received")
			var event types.Event
			err := json.Unmarshal([]byte(*message.Body), &event)
			if err != nil {
				fmt.Printf("Error unmarshalling message: %v\n", err)
				return
			}
			fmt.Println(event, "IS THE EVENT")
			if event.Type == chatserversqs.MESSAGE_GROUP_SERVER {
				ServerQueries.Content = append(ServerQueries.Content, event.Content)
				ServerQueries.Type = append(ServerQueries.Type, event.MessageType)
				ServerQueries.ReplyOf = append(ServerQueries.ReplyOf, event.ReplyOf)
				ServerQueries.TimeAt = append(ServerQueries.TimeAt, event.TimeAt)
				ServerQueries.ChannelID = append(ServerQueries.ChannelID, event.ChannelId)
				ServerQueries.FromUserID = append(ServerQueries.FromUserID, event.FromID)
			} else if event.Type == chatserversqs.MESSAGE_GROUP_USER {
				UserQueries.Content = append(UserQueries.Content, event.Content)
				UserQueries.Type = append(UserQueries.Type, event.MessageType)
				UserQueries.ReplyOf = append(UserQueries.ReplyOf, event.ReplyOf)
				UserQueries.TimeAt = append(UserQueries.TimeAt, event.TimeAt)
				UserQueries.ToUserID = append(UserQueries.ToUserID, event.To)
				UserQueries.FromUserID = append(UserQueries.FromUserID, event.FromID)
			} else {

			}
			_, err = chatserversqs.SQSClient.DeleteMessage(ctx, &sqs.DeleteMessageInput{
				QueueUrl:      &chatserversqs.SQSQueueURL,
				ReceiptHandle: message.ReceiptHandle,
			})
		}
		fmt.Println("UserQueries:", UserQueries)
		fmt.Println("ServerQueries:", ServerQueries)

		if len(UserQueries.Content) > 0 {
			err := db.DBQueries.BatchInsertUserToUserChat(ctx, UserQueries)
			if err != nil {
				fmt.Println("Error inserting user messages:", err.Error())
			}
		}

		if len(ServerQueries.Content) > 0 {
			err := db.DBQueries.BatchInsertUserToChannelChat(ctx, ServerQueries)
			if err != nil {
				fmt.Println("Error inserting server messages:", err.Error())
			}
		}
		UserQueries = db.BatchInsertUserToUserChatParams{
			Content:    make([]string, 0),
			FromUserID: make([]int64, 0),
			ToUserID:   make([]int64, 0),
			Type:       make([]string, 0),
			TimeAt:     make([]int64, 0),
		}
		ServerQueries = db.BatchInsertUserToChannelChatParams{
			Content:    make([]string, 0),
			ReplyOf:    make([]int64, 0),
			FromUserID: make([]int64, 0),
			ChannelID:  make([]int64, 0),
			Type:       make([]string, 0),
			TimeAt:     make([]int64, 0),
		}
		time.Sleep(PERSIST_FREQUENCY)
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file", err.Error())
	}
	chatserversqs.InitSQSClient()
	PersistWorker()
}
