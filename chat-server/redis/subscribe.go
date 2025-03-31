package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func Ticker(ctx context.Context, sub *redis.PubSub, channel string) {
	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			active, _ := IsTopicActive(ctx, channel)
			if !active {
				fmt.Println("Topic expired, unsubscribing:", channel)
				sub.Close()
				return
			}
		case <-ctx.Done():
			fmt.Println("Context canceled, unsubscribing:", channel)
			sub.Close()
			return
		}

	}
}
func Subscribe(ctx context.Context, userId int64, channels []string, callback func(string, int64)) {
	if RedisClient == nil {
		return
	}
	sub := RedisClient.Subscribe(ctx, channels...)
	ch := sub.Channel()
	fmt.Println("Subscribed to channel:", channels)
	for msg := range ch {
		callback(msg.Payload, userId)
	}
}
