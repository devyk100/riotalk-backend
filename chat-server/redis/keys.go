package redis

import "strconv"

func ServerKey(channelId int64) string {
	return "server-" + strconv.FormatInt(channelId, 10)
}

func UserKey(channelId int64) string {
	return "user-" + strconv.FormatInt(channelId, 10)
}
