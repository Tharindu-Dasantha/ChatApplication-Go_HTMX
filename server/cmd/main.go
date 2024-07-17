package main

import (
	"log"
	"github.com/gorilla/websocket"
	"net/http"
)


var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func main() {

}