package main

type Hub struct {
	clients    map[*client]bool
	broadcast  chan []byte
	register   chan *client
	unregister chan *client
}

var hub = &Hub{clients: make(map[*client]bool)}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case bytes := <-h.broadcast:
			for client := range h.clients {
				client.send <- bytes
			}
		}
	}
}