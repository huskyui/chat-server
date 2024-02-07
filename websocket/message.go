package websocket

type Message struct {
	Type    string `json:"type"`
	Content string `json:"content"`
	User    string `json:"user"`
}
