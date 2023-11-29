package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

type client struct {
	conn *websocket.Conn
	send chan []byte
	hub  *Hub
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
	//go client.writePump()
	for {
		_, message, err := conn.ReadMessage()
		fmt.Println("readmessage", message)
		if err != nil {
			panic(err)
		}
		hub.broadcast <- message
	}

}

func (c *client) readPump() {
	defer func() {
		fmt.Println("end readPump")
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		fmt.Println("readmessage", message)
		if err != nil {
			panic(err)
		}
		c.hub.broadcast <- message
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
