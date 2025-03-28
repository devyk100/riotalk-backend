package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var GoogleTokenUrl = "https://oauth2.googleapis.com/token"
var GoogleUserInfoUrl = "https://www.googleapis.com/oauth2/v2/userinfo"

type GoogleTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	IDToken      string `json:"id_token"`
}

func GoogleCallback() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Query("code")
		fmt.Println("Code: ", code)
		data := url.Values{}
		data.Set("code", code)
		data.Set("client_id", os.Getenv("CLIENT_ID"))
		data.Set("client_secret", os.Getenv("CLIENT_SECRET"))
		data.Set("redirect_uri", CallbackUrl)
		data.Set("grant_type", "authorization_code")
		req, err := http.NewRequest("POST", GoogleTokenUrl, strings.NewReader(data.Encode()))
		if err != nil {
			log.Fatalf("Error creating request: %v", err)
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("Error making POST request: %v", err)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Error reading response body: %v", err)
		}

		responseStr := string(body)
		fmt.Println("Response Status:", resp.Status)
		fmt.Println("Response Body:", responseStr)

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Println("Error closing response body", err.Error())
			}
		}(resp.Body)
		var googleTokenResponse GoogleTokenResponse
		err = json.Unmarshal(body, &googleTokenResponse)
		if err != nil {
			log.Fatalf("Error decoding JSON: %v", err)
		}

		c.SetCookie("refresh_token", googleTokenResponse.RefreshToken, 60*60*24, "/", "", false, true)
		c.Redirect(http.StatusFound, "http://localhost:3000/auth-success")
	}
}

func GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

type RefreshTokenRequest struct {
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

func RefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshToken, err := c.Cookie("refresh_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No refresh token"})
			return
		}

		reqBody, _ := json.Marshal(RefreshTokenRequest{
			ClientID:     os.Getenv("CLIENT_ID"),
			ClientSecret: os.Getenv("CLIENT_SECRET"),
			RefreshToken: refreshToken,
			GrantType:    "refresh_token",
		})

		resp, err := http.Post(GoogleTokenUrl, "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to contact Google"})
			return
		}
		body, err := io.ReadAll(resp.Body)
		defer resp.Body.Close()
		fmt.Println(string(body))
		var tokenResp RefreshTokenResponse
		if err := json.Unmarshal(body, &tokenResp); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"accessToken": tokenResp.AccessToken})
	}
}

func InitiateGoogleAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		client_id := os.Getenv("CLIENT_ID")
		c.Redirect(http.StatusFound, "https://accounts.google.com/o/oauth2/auth?client_id="+client_id+
			"&redirect_uri="+CallbackUrl+
			"&response_type=code&scope=email%20profile&access_type=offline&prompt=consent")
	}
}

var CallbackUrl = "http://localhost:8080/auth/callback/google"

func GetOauthURL() gin.HandlerFunc {
	return func(c *gin.Context) {
		client_id := os.Getenv("CLIENT_ID")
		c.JSON(200, gin.H{
			"url": "https://accounts.google.com/o/oauth2/auth?client_id=" + client_id +
				"&redirect_uri=" + CallbackUrl +
				"&response_type=code&scope=email%20profile&access_type=offline&prompt=consent",
		})
	}
}

func AuthRouter(router *gin.Engine) *gin.RouterGroup {
	authRouter := router.Group("/auth")
	authRouter.GET("/get-oauth-url", GetOauthURL())
	authRouter.GET("/callback/google", GoogleCallback())
	authRouter.GET("/initiate/google", InitiateGoogleAuth())
	authRouter.GET("/refresh-token/google", RefreshToken())
	authRouter.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("refresh_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No refresh token"})
		} else {
			c.String(http.StatusOK, cookie)
		}
	})
	return authRouter
}
