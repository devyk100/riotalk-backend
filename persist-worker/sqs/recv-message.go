package sqs

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	sqs_types "github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"log"
)

func ReceiveMessage() ([]sqs_types.Message, error) {
	output, err := SQSClient.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
		QueueUrl:            &SQSQueueURL,
		MaxNumberOfMessages: 10, // Get up to 10 messages (max)
		WaitTimeSeconds:     20, // Blocks for up to 20 seconds -> max unsee time for the other ones
	})
	if err != nil {
		log.Printf("Error receiving messages: %v\n", err)
		return nil, err
	}
	return output.Messages, nil
}
