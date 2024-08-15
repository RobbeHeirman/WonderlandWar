package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	socketManager := NewManager()
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		socketManager.NewConnection(writer, request)
	})
	http.ListenAndServe(":8080", nil)
}
