package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func LoadPageAndMethodsWebSockets(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("Web-Socket")
		Handler(w, r)
	}
}

func Handler(writer http.ResponseWriter, request *http.Request) string {
	socket, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println(err)
		return "Upgrader error"
	}
	for {
		// Reading WebSocket Message
		msgType, msg, err := socket.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return "Message Received Error"
		}
		// Logging
		fmt.Println("Message received: ", string(msg))

		// Returning Message
		err = socket.WriteMessage(msgType, msg)
		if err != nil {
			fmt.Println(err)
			return "Error Returning Message"
		}
	}
}
