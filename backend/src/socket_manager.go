package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/robbeheirman/wonderland-war/collections"
	"github.com/robbeheirman/wonderland-war/proto_messages"
)

func NewManager() *SocketManager {
	return &SocketManager{
		connections: collections.NewUnOrderedList[PlayerSocketController](),
	}
}

type SocketManager struct {
	connections collections.ConcurrentUnorderedList[PlayerSocketController]
}

func (manager *SocketManager) NewConnection(w http.ResponseWriter, r *http.Request) {
	connection, _ := upgrader.Upgrade(w, r, nil)

	for {
		msgtype, message, err := connection.ReadMessage()
		if err != nil || msgtype == websocket.CloseMessage {
			fmt.Println("Connection closed")
			break
		}

		envelope := proto_messages.Envelope{}
		if err := proto.Unmarshal(message, &envelope); err != nil {
			log.Fatalln("Could not unmarshal envelope", err)
		}

		m, err := envelope.Data.UnmarshalNew()
		switch _ := m.(type) {
		case *proto_messages.JoinLobbyMessage:
		}
	}
}

func (manager *SocketManager) joinLobby(msg *proto_messages.JoinLobbyMessage, conn *websocket.Conn) {
	anyMsg, err := anypb.New(msg)
	if err != nil {
		log.Fatalln("Could not marshal anypb", err)
	}
	envelope := proto_messages.Envelope{
		Data: anyMsg,
	}
	manager.Broadcast(&envelope)
	ctrl := PlayerSocketController{
		PlayerModel: &PlayerModel{},
		Connection:  conn,
	}
	manager.connections.Append(&ctrl)
}

func (manager *SocketManager) Broadcast(msg *proto_messages.Envelope) {
	marshalled, err := proto.Marshal(msg)
	if err != nil {
		log.Fatalln("Could not marshal envelope", err)
	}
	manager.connections.Apply(func(w *PlayerSocketController) {
		err := w.Connection.WriteMessage(websocket.BinaryMessage, marshalled)
		if err != nil {
			log.Fatalln("Could not write envelope", err)
		}
	})
}
