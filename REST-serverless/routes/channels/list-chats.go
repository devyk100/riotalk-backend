package channels_route

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shared/redis"
	"strconv"
)

type ChatMessage struct {
	ChannelID    int64  `json:"channel_id"`
	Content      string `json:"content"`
	Event        string `json:"event"`
	FromID       int64  `json:"from_id"`
	FromImg      string `json:"from_img"`
	FromName     string `json:"from_name"`
	FromUsername string `json:"from_username"`
	MessageType  string `json:"message_type"`
	ReplyOf      int64  `json:"reply_of"`
	TimeAt       int64  `json:"time_at"`
	To           int64  `json:"to"`
	Type         string `json:"type"`
}

func ListChats() gin.HandlerFunc {
	return func(c *gin.Context) {
		serverIDStr := c.Query("server_id")
		channelIDStr := c.Query("channel_id")

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

		rawMessages, err := redis.GetRecentMessages(c.Request.Context(), redis.RecentMessageServerKey(channelID, serverID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch messages"})
			return
		}

		var chatMessages []ChatMessage
		for _, msgStr := range rawMessages {
			var msg ChatMessage
			if err := json.Unmarshal([]byte(msgStr), &msg); err == nil {
				chatMessages = append(chatMessages, msg)
			} else {
				// log parse error and skip bad message
				fmt.Println("Failed to parse message from Redis:", err)
			}
		}

		c.JSON(http.StatusOK, chatMessages)
	}
}
