package main

import (
	"chat-server/cmd"
	"chat-server/redis"
	"chat-server/sqs"
	"chat-server/types"
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"net/http"
	"sync"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // (Change this in production)
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer func() {
		conn.Close()

	}()
	userId, err := cmd.HandleWSAuth(conn)
	Client := &types.Client{
		Conn: conn,
		Mu:   &sync.Mutex{},
	}
	if err != nil {
		fmt.Println("Error handling WSAuth:", err)
		return
	}

	// DO NOT EXIT THIS FUNCTION, THE OTHER CREATED CONN'S, AND OTHER'S REFERENCES WOULD BE LOST, AND DEREFERENCING ERRORS WILL OCCUR
	fmt.Println("New client connected!")
	go cmd.MainEventLoop(Client, userId)
	cmd.MainReadLoop(Client, userId)
}

func main() {
	http.HandleFunc("/ws", handleConnections)
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file", err.Error())
	}
	port := "8090"
	fmt.Println("WebSocket server started on : " + port)
	err = redis.InitRedisClient(context.Background())
	if err != nil {
		fmt.Println("Error initializing Redis client:", err)
		return
	}
	sqs.InitSQSClient()
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
