package redis

import (
	"context"
)

func Publish(ctx context.Context, channel string, message string) error {
	err := RedisClient.Publish(ctx, channel, message).Err()
	if err != nil {
		return err
	}
	return nil
}
