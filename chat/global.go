package chat

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	Clients = make(map[string]*Client)
	mu      sync.Mutex
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}
