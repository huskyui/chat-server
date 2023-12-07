package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type client struct {
	conn *websocket.Conn
	send chan []byte
	hub  *Hub
	name string
}

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  512,
	WriteBufferSize: 512,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func serverWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	client := &client{hub: hub, conn: conn, send: make(chan []byte)}
	hub.register <- client
	go client.writePump()
	go client.readPump()

}

func (c *client) readPump() {
	defer func() {
		fmt.Println("end readPump")
	}()

	for {
		_, messageBytes, err := c.conn.ReadMessage()
		if err != nil {
			c.conn.Close()
			break
		}
		var message Message
		e := json.Unmarshal(messageBytes, &message)
		if e != nil {
			log.Fatal(err)
		}
		switch message.Type {
		case "login":
			c.name = message.Content
		case "broadcast":
			c.hub.broadcast <- messageBytes
		}
	}
}

func (c *client) writePump() {
	defer func() {
		print("end writePump")
	}()
	for {
		select {
		case bytes, ok := <-c.send:
			if !ok {
				return
			}
			c.conn.WriteMessage(websocket.TextMessage, bytes)
		}
	}
}
