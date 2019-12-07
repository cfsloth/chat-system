package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	Message     string
	Sender      string
	Receiver    string
	MessageType int
	Error       bool
}

var message *Message = new(Message)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func LoadPageAndMethodsWebSockets(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("Web-Socket")
		fmt.Println(Handler(w, r))
	}
}

/* free message */
func Handler(writer http.ResponseWriter, request *http.Request) *Message {
	socket, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println(err)
		message.Message = "Upgrade Error"
		message.Receiver = "Receiver"
		message.Sender = "Sender"
		message.Error = true
		return message
	}
	for {
		// Reading WebSocket Message
		msgType, msg, err := socket.ReadMessage()
		if err != nil {
			fmt.Println(err)
			message.Message = "Receiving Error"
			message.Receiver = "Request"
			message.Sender = "Sender"
			message.MessageType = msgType
			message.Error = true
			return message
		}
		// Logging
		fmt.Println("Message received: ", string(msg))
		message.Message = "Upgrade Error"
		message.Receiver = "Request"
		message.Sender = "Sender"
		message.MessageType = msgType
		message.Error = false
		err = socket.WriteMessage(msgType, msg)
		if err != nil {
			fmt.Println(err)
			return message
		}
	}
}
