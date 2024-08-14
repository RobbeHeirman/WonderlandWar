package main

import (
	"WonderlandWar/proto_messages"
	"fmt"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func RunSocket(w http.ResponseWriter, r *http.Request) {
	connection, _ := upgrader.Upgrade(w, r, nil)
	for {
		msgtype, message, err := connection.ReadMessage()
		if err != nil || msgtype == websocket.CloseMessage {
			fmt.Println("Connection closed")
			break
		}
		fmt.Printf("Received message: %s\n", string(message))
	}

}

func main() {
	t := &proto_messages.JoinLobbyMessage{
		Name: "Test",
	}
	out, _ := proto.Marshal(t)
	fmt.Println(out)
	http.HandleFunc("/", RunSocket)
	http.ListenAndServe(":8080", nil)
}
