package helper

import (
	"fmt"
	"math/rand"
	"time"
)

var s1 = rand.NewSource(time.Now().UnixNano())
var r1 = rand.New(s1)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client := range pool.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})
			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Szie of Connection Pool: ", len(pool.Clients))
			for client := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
			}
			break
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in Pool")
			for client := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}

func Emit(pool *Pool) {
	for {
		select {
		default:
			i := r1.Intn(len(usernameList))
			j := r1.Intn(len(messageList))

			time.Sleep(4 * time.Second)
			for client := range pool.Clients {
				client.Conn.WriteJSON(Emission{User: usernameList[i], Message: messageList[j]})
			}
			fmt.Printf("Broadcasted message: {User: %s, Message: %s}\n", usernameList[i], messageList[j])

		}
	}
}
