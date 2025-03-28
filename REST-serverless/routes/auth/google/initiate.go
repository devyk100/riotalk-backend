package google

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func InitiateGoogleAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		client_id := os.Getenv("CLIENT_ID")
		c.Redirect(http.StatusFound, "https://accounts.google.com/o/oauth2/auth?client_id="+client_id+
			"&redirect_uri="+CallbackUrl+
			"&response_type=code&scope=email%20profile&access_type=offline&prompt=consent")
	}
}
