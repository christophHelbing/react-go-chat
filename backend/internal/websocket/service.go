package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Service interface {
	RegisterNewClient(pool *Pool, writer http.ResponseWriter, request *http.Request)
}

type websocketFactory struct {
	upgrader websocket.Upgrader
}

func NewInstance() Service {
	return &websocketFactory{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (factory websocketFactory) RegisterNewClient(pool *Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	conn, err := factory.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	client := createClient(conn, pool)

	pool.Register <- client
	client.Read()
}

func createClient(conn *websocket.Conn, pool *Pool) *Client {
	return &Client{
		Connection: conn,
		Pool:       pool,
	}
}
