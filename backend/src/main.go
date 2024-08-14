package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/robbeheirman/wonderland-war/proto_messages"
	"google.golang.org/protobuf/proto"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func RunSocket(w http.ResponseWriter, r *http.Request) {

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
