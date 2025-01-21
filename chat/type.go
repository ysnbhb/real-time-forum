package chat

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Id   int
	Name string
	Conn websocket.Conn
}

// func (conn *Client) Read() {
// 	conn.Conn.ReadJSON()
// }

// func (conn *Client) Read() {
// 	conn.Conn.WriteJSON()
// }


