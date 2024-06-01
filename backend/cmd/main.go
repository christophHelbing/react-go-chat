package main

import (
	"fmt"
	"github.com/christophHelbing/react-go-chat/internal/websocket"
	"net/http"
)

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		websocket.NewInstance().RegisterNewClient(pool, writer, request)
	})
}
func main() {
	fmt.Println("Distributed Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
