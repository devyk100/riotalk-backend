package redis

import (
	"context"
	"fmt"
)

func Subscribe(ctx context.Context, channel string, callback func(string)) {
	if RedisClient == nil {
		return
	}
	sub := RedisClient.Subscribe(ctx, channel)
	ch := sub.Channel()
	fmt.Println("Subscribed to channel:", channel)
	for msg := range ch {
		callback(msg.Payload)
	}
}
