package chat

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var Clients = make(map[string]*Client)

var UpGrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}
