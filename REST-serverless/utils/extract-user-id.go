package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func ExtractUserIDFromContext(c *gin.Context) (int64, error) {
	userId, exists := c.Get("userId")
	if !exists {
		return -1, errors.New("user Id not found in context")
	}
	userIdInt64, ok := userId.(int64)
	if !ok {
		return -1, errors.New("user Id not found in context")
	}
	return userIdInt64, nil
}
