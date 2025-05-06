package redis

import (
	"context"
	"fmt"
)

func GetRecentMessages(ctx context.Context, channelID string) ([]string, error) {
	messages, err := RedisClient.LRange(ctx, channelID, 0, -1).Result()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return messages, nil
}
