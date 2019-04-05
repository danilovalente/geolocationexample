package usecase

import (
	"github.com/danilovalente/geolocationexample/domain"
	"github.com/danilovalente/geolocationexample/gateway/appwebsocket"
	"github.com/danilovalente/geolocationexample/gateway/mongodb"
)

//GetTransports returns all the transports
func GetTransports() ([]*domain.Transport, error) {
	return mongodb.TransportRepo.GetAll()
}

//GetTransport returns a specific transport by ID
func GetTransport(id string) (*domain.Transport, error) {
	return mongodb.TransportRepo.Get(id)
}

//CreateTransport creates a new Transport in the database.
func CreateTransport(transport *domain.Transport) (string, error) {
	return mongodb.TransportRepo.Save(transport)
}

//UpdateTransportPosition updates a Transports current position and notifies the clients listening.
func UpdateTransportPosition(transportID string, position *domain.Position) (*domain.Transport, error) {
	transport, err := mongodb.TransportRepo.UpdatePosition(transportID, position)
	if err != nil {
		return nil, err
	}
	notification := domain.PositionChangeNotificationFromTransport(transport)
	appwebsocket.WebsocketServer.Broadcast(notification)

	return transport, nil
}
