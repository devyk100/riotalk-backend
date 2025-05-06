package redis

import "strconv"

func ServerKey(channelId int64) string {
	return "server-" + strconv.FormatInt(channelId, 10)
}

func UserKey(channelId int64) string {
	return "user-" + strconv.FormatInt(channelId, 10)
}

func RecentMessageServerKey(channelId int64, serverId int64) string {
	return "recent-" + ServerKey(serverId) + "-" + strconv.FormatInt(channelId, 10)
}

func RecentMessageUserKey(To int64, By int64) string {
	return "recent-" + UserKey(min(To, By)) + "-" + strconv.FormatInt(max(To, By), 10)
}
