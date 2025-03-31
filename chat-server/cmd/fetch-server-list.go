package cmd

import (
	"chat-server/state"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ServerData struct {
	ID          int64  `json:"ServerID"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

func FetchServerList(userID int64) ([]ServerData, error) {
	url := "http://localhost:8080/servers/list"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("MAKING A REQUEST TO ", url, "with", state.AccessTokens[userID])
	req.Header.Set("Authorization", "Bearer "+state.AccessTokens[userID])
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch servers: %s", resp.Status)
	}

	var servers []ServerData
	err = json.Unmarshal(body, &servers)
	if err != nil {
		return nil, err
	}

	return servers, nil
}
