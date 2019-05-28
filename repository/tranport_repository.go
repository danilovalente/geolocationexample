package repository

import "github.com/danilovalente/geolocationexample/domain"

//TransportRepositoryName contains the name of the TransportRepository in the RepoMap
const TransportRepositoryName = "TransportRepository"

/*
TransportRepository defines the repository capabilities that should be found in a Repository implementation for Transport
*/
type TransportRepository interface {
	Repository
	Get(id string) (*domain.Transport, error)
	GetAll() ([]*domain.Transport, error)
	Save(transport *domain.Transport) (string, error)
	UpdatePosition(transportID string, newPosition *domain.Position) (*domain.Transport, error)
}

//GetTransportRepository gets the TransportRepository current implementation
func GetTransportRepository() TransportRepository {
	return Repos.Get(TransportRepositoryName).(TransportRepository)
}
