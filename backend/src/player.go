package main

import (
	"github.com/gorilla/websocket"
)

type PlayerModel struct {
	Id   int
	Name string
}

type PlayerSocketController struct {
	PlayerModel *PlayerModel
	Connection  *websocket.Conn
}
