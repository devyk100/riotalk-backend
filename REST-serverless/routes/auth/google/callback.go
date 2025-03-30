package google

import (
	"REST-serverless/db"
	"REST-serverless/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

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
		// get the code from when google calls back this url
		code := c.Query("code")

		// set the payload, to send them as x www form encoded, not json
		data := url.Values{}
		data.Set("code", code)
		data.Set("client_id", os.Getenv("CLIENT_ID"))
		data.Set("client_secret", os.Getenv("CLIENT_SECRET"))
		data.Set("redirect_uri", CallbackUrl)
		data.Set("grant_type", "authorization_code")

		// get ready for the request command
		req, err := http.NewRequest("POST", GoogleTokenUrl, strings.NewReader(data.Encode()))
		if err != nil {
			log.Fatalf("Error creating request: %v", err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		client := &http.Client{}
		// execute the request
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("Error making POST request: %v", err)
		}

		// process the body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Error reading response body: %v", err)
		}
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

		// fetch the user data from the Google's backend
		user, err := FetchGoogleUserData(&googleTokenResponse.AccessToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		fmt.Println(user)
		// make a DB request to check if the User was available, or not
		userData, err := db.DBQueries.CreateUserOrDoNothing(c.Request.Context(), db.CreateUserOrDoNothingParams{
			Name:     user.Name,
			Username: strings.Fields(user.Name)[0] + utils.RandomString(15),
			Email:    user.Email,
			Img: pgtype.Text{
				String: user.Picture,
				Valid:  true,
			},
			Description: pgtype.Text{
				String: "",
				Valid:  false,
			},
			Provider: "google",
			Verified: true,
			Password: pgtype.Text{
				Valid: false,
			},
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		method := "google"
		token := utils.CreateRefreshToken(method, googleTokenResponse.RefreshToken, userData.ID)
		fmt.Println(token, "is the token, generated")
		c.SetCookie("refresh_token", token, 60*60*24, "/", "", false, true)
		c.Redirect(http.StatusFound, "http://localhost:3000/auth-success")
	}
}
