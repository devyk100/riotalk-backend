package redis

import (
	"context"
	"fmt"
)

var RECENT_MESSAGE_THRESHOLD = int64(25)

func PushToRecentMessages(ctx context.Context, channel string, message string) {
	err := RedisClient.RPush(ctx, channel, message).Err()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = RedisClient.LTrim(ctx, channel, -RECENT_MESSAGE_THRESHOLD, -1).Err()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
