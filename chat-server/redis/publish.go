package redis

import (
	"context"
	"fmt"
	"time"
)

func Publish(ctx context.Context, channel string, message string) error {
	RedisClient.Set(ctx, channel, true, time.Minute*10)
	err := RedisClient.Publish(ctx, channel, message).Err()
	if err != nil {
		return err
	} else {
		fmt.Println("Published:", message)
	}
	return nil
}
