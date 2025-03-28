package google

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GoogleUser struct {
	Sub           string `json:"sub"` // Unique ID
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

func FetchGoogleUserData(accessToken *string) (*GoogleUser, error) {
	// get user id from database
	req, err := http.NewRequest("GET", GoogleUserInfoUrl, nil)
	if err != nil {
		//return nil, err
	}

	// Set the Authorization header
	req.Header.Set("Authorization", "Bearer "+*accessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body", err.Error())
		}
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch user info, status: %d", resp.StatusCode)
	}

	var user GoogleUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
