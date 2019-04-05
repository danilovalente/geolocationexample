package controller

import (
	"fmt"
	"net/http"

	"github.com/danilovalente/geolocationexample/gateway/appwebsocket"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func websocketConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	appwebsocket.WebsocketServer.AddClientConnection(conn)
	fmt.Println("Client subscribed")
}
