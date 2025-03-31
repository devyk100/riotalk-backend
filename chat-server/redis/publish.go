package redis

import (
	"context"
	"fmt"
)

func Publish(ctx context.Context, channel string, message string) error {
	err := RedisClient.Publish(ctx, channel, message).Err()
	if err != nil {
		return err
	} else {
		fmt.Println("Published:", message)
	}
	return nil
}
