package google

import (
	"REST-serverless/utils"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

type RefreshTokenGoogleRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RefreshToken string `json:"refresh_token"`
	GrantType    string `json:"grant_type"`
}

type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func RefreshTokenGoogle(refreshToken string, c *gin.Context) {
	reqBody, err := json.Marshal(RefreshTokenGoogleRequest{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RefreshToken: refreshToken,
		GrantType:    "refresh_token",
	})
	if err != nil {

	}

	resp, err := http.Post(GoogleTokenUrl, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to contact Google"})
		return
	}
	body, err := io.ReadAll(resp.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var tokenResp RefreshTokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response"})
		return
	}

	method := "google"
	token, err := utils.CreateAccessToken(&method, &tokenResp.AccessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"accessToken": token})
}
