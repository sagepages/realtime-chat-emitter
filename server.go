package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	h "github.com/sagepages/emitter/helper"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Allowing an origin for testing
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func websocketHandler(pool *h.Pool, w http.ResponseWriter, r *http.Request) {

	// Handle upgrading the connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error in handler:", err)
		return
	}

	client := &h.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()

	log.Println("Client connected.")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func setupRoutes() {
	pool := h.NewPool()

	go pool.Start()
	go h.Emit(pool)

	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocketHandler(pool, w, r)
	})
}

func main() {
	fmt.Println("Server started")
	h.GenerateLists()
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
