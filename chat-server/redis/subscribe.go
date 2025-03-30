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
func Subscribe(ctx context.Context, channel string, callback func(string)) {
	active, err := IsTopicActive(ctx, channel)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if RedisClient == nil || !active {
		return
	}
	sub := RedisClient.Subscribe(ctx, channel)
	ch := sub.Channel()
	go Ticker(ctx, sub, channel)
	fmt.Println("Subscribed to channel:", channel)
	for msg := range ch {
		callback(msg.Payload)
	}
}
