package redis

import (
	"context"
	"log"
)

func IsTopicActive(ctx context.Context, channel string) (bool, error) {
	exists, err := RedisClient.Exists(ctx, channel).Result()
	if err != nil {
		log.Println("Error checking key:", err)
		return false, err
	}
	if exists == 0 {
		return false, nil
	} else {
		return true, nil
	}
}
