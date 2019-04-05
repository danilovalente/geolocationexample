package appwebsocket

import (
	"encoding/json"
	"log"

	"github.com/danilovalente/geolocationexample/domain"
	"github.com/gorilla/websocket"
)

var WebsocketServer WebsocketManager

type WebsocketManager struct {
	websocketClients []*websocket.Conn
}

//AddClientConnection to client connection catalog for broadcasting
func (websocketManager *WebsocketManager) AddClientConnection(conn *websocket.Conn) {
	websocketManager.websocketClients = append(websocketManager.websocketClients, conn)
}

func (websocketManager *WebsocketManager) RemoveClientConnection(i int) {
	copy(websocketManager.websocketClients[i:], websocketManager.websocketClients[i+1:])
	websocketManager.websocketClients[len(websocketManager.websocketClients)-1] = nil
	websocketManager.websocketClients = websocketManager.websocketClients[:len(websocketManager.websocketClients)-1]
}

//Broadcast sends a notification of positioning changes to all clients connected
func (websocketManager *WebsocketManager) Broadcast(positionChangeNotification *domain.PositionChangeNotification) {
	for i := 0; i < len(websocketManager.websocketClients); i++ {
		clientConn := websocketManager.websocketClients[i]
		positionChangeNotificationJSON, err := json.Marshal(positionChangeNotification)
		if err != nil {
			log.Println("An error occurred while trying to serialize a positionChangeNotification")
		} else {
			log.Println("Broadcasting position change: " + string(positionChangeNotificationJSON))
			err := clientConn.WriteMessage(websocket.TextMessage, positionChangeNotificationJSON)
			if err != nil {
				log.Println("An error occurred while trying to broadcast a positionChangeNotification")
				clientConn.Close()
				websocketManager.RemoveClientConnection(i)
			}
		}
	}
}

func init() {
	WebsocketServer = WebsocketManager{websocketClients: make([]*websocket.Conn, 0)}
}
