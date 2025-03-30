package main

import (
	"chat-server/cmd"
	"chat-server/redis"
	"chat-server/state"
	"chat-server/types"
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"net/http"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections (Change this in production)
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()
	userId, err := cmd.HandleWSAuth(conn)
	state.Clients[userId] = conn
	state.Events[userId] = make(chan *types.Event)
	if err != nil {
		fmt.Println("Error handling WSAuth:", err)
		return
	}

	// DO NOT EXIT THIS FUNCTION, THE OTHER CREATED CONN'S, AND OTHER'S REFERENCES WOULD BE LOST, AND DEREFERENCING ERRORS WILL OCCUR
	fmt.Println("New client connected!")
	go cmd.MainWriteLoop(conn, userId)
	cmd.MainReadLoop(conn, userId)
}

func main() {
	// Setup HTTP server
	http.HandleFunc("/ws", handleConnections)
	err := godotenv.Load(".env")
	port := "8090"
	fmt.Println("WebSocket server started on : " + port)
	err = redis.InitRedisClient(context.Background())
	if err != nil {
		fmt.Println("Error initializing Redis client:", err)
		return
	}
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
