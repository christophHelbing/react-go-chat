package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type WebsocketService interface {
	ServeWebSocket(writer http.ResponseWriter, request *http.Request)
}

type websocketFactory struct {
	upgrader websocket.Upgrader
}

func NewInstance() WebsocketService {
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

func (factory websocketFactory) ServeWebSocket(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Host)

	ws, err := factory.upgrader.Upgrade(writer, request, nil)

	if err != nil {
		log.Println(err)
	}

	reader(ws)
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}
