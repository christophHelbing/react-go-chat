package main

import (
	"fmt"
	"github.com/christophHelbing/react-go-chat/internal/websocket"
	"net/http"
)

func setupRoutes() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Simple Server")
	})
	http.HandleFunc("/ws", websocket.NewInstance().ServeWebSocket)
}
func main() {
	fmt.Println("Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
