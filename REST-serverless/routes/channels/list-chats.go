package channels_route

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shared/db"
	"shared/redis"
	"shared/types"
	"strconv"
)

func ListChats() gin.HandlerFunc {
	return func(c *gin.Context) {
		serverIDStr := c.Query("server_id")
		channelIDStr := c.Query("channel_id")
		beforeStr := c.DefaultQuery("before", "-1")

		serverID, err := strconv.ParseInt(serverIDStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid server_id"})
			return
		}
		channelID, err := strconv.ParseInt(channelIDStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid channel_id"})
			return
		}

		before, err := strconv.ParseInt(beforeStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid before"})
		}

		if before == -1 {
			rawMessages, err := redis.GetRecentMessages(c.Request.Context(), redis.RecentMessageServerKey(channelID, serverID))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch messages"})
				return
			}

			var chatMessages []types.Event
			for _, msgStr := range rawMessages {
				var msg types.Event
				if err := json.Unmarshal([]byte(msgStr), &msg); err == nil {
					chatMessages = append(chatMessages, msg)
				} else {
					// log parse error and skip bad message
					fmt.Println("Failed to parse message from Redis:", err)
				}
			}

			c.JSON(http.StatusOK, chatMessages)
		} else {
			rawMessages, err := db.DBQueries.GetChannelChatsBefore(c.Request.Context(), db.GetChannelChatsBeforeParams{ChannelID: channelID, TimeAt: before})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch messages"})
				return
			}

			var chatMessages []types.Event
			for _, msg := range rawMessages {
				var chatMsg types.Event
				chatMsg.ChannelId = msg.ChannelID
				chatMsg.TimeAt = msg.TimeAt
				chatMsg.Content = msg.Content.String
				chatMsg.To = msg.ServerID
				chatMsg.Event = "chat"
				chatMsg.FromID = msg.FromUserID
				chatMsg.FromImg = msg.UserImg.String
				chatMsg.FromName = msg.UserName
				chatMsg.FromUsername = msg.UserUsername
				chatMessages = append(chatMessages, chatMsg)
			}

			c.JSON(http.StatusOK, chatMessages)
		}
	}
}
