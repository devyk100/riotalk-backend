package redis

import (
	"context"
	"fmt"
	"sort"
)

func GetRecentMessages(ctx context.Context, channelID string) ([]string, error) {
	messages, err := RedisClient.LRange(ctx, channelID, 0, RECENT_MESSAGE_THRESHOLD).Result()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	sort.Slice(messages, func(i, j int) bool { return i > j })
	return messages, nil
}
