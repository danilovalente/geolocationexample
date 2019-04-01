package usecase

import (
	"github.com/danilovalente/geolocationexample/domain"
	"github.com/danilovalente/geolocationexample/gateway/mongodb"
)

//GetTransports returns the transports
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

//UpdateTransportPosition updates a Transports current position.
func UpdateTransportPosition(transportID string, position *domain.Position) (*domain.Transport, error) {
	return mongodb.TransportRepo.UpdatePosition(transportID, position)
}
