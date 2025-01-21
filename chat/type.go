package chat

import (
	"net/http"

	"real-time-forum/utils"

	"github.com/gorilla/websocket"
)

type Client struct {
	Id   int
	Name string
	Conn *websocket.Conn
	Uid  string
}

type Api struct {
	Cn Contat
}

type Contat interface {
	GetIdFromReq(*http.Request) (int, string, string)
}

func (conn *Client) Read(message *utils.Message) error {
	return conn.Conn.ReadJSON(message)
}

func (conn *Client) Write(message any) {
	err := conn.Conn.WriteJSON(message)
	if err != nil {
		conn.Conn.Close()
	}
}

func (api *Api) Chat(w http.ResponseWriter, r *http.Request) {
	id, name, uid := api.Cn.GetIdFromReq(r)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to upgrade WebSocket", http.StatusInternalServerError)
		return
	}

	client := NewComm(id, name, conn, uid)
	mu.Lock()
	Clients[uid] = client
	mu.Unlock()

	go api.ReadAndWrite(client)
}

func (api *Api) ReadAndWrite(client *Client) {
	message := utils.Message{}
	defer func() {
		client.Conn.Close()
		mu.Lock()
		delete(Clients, client.Uid)
		mu.Unlock()
	}()

	for {
		err := client.Read(&message)
		if err != nil {
			break
		}
		mu.Lock()
		conn, exists := Clients[message.To]
		mu.Unlock()
		if exists {
			conn.Write(message)
		}
	}
}

func NewComm(id int, name string, conn *websocket.Conn, uid string) *Client {
	return &Client{
		Id:   id,
		Name: name,
		Conn: conn,
		Uid:  uid,
	}
}
