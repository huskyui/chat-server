package websocket

import (
	"encoding/json"
)

type Hub struct {
	clients    map[*client]bool
	broadcast  chan []byte
	register   chan *client
	unregister chan *client
	login      chan *client
}

var hub = &Hub{
	clients:    make(map[*client]bool),
	broadcast:  make(chan []byte),
	register:   make(chan *client),
	unregister: make(chan *client),
	login:      make(chan *client),
}

func WebsocketHub() *Hub {
	return hub
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			sendHandShake(client)
		case bytes := <-h.broadcast:
			for client := range h.clients {
				client.send <- bytes
			}
		}
	}
}

func sendHandShake(c *client) {
	message := Message{Type: "handshake", Content: "hello from server"}
	bytes, _ := json.Marshal(message)
	c.send <- bytes
}
