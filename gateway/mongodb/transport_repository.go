package mongodb

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/danilovalente/geolocationexample/domain"
	"github.com/danilovalente/geolocationexample/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//DatabaseName in MongoDB
const DatabaseName = "geolocation"

//CollectionName in MongoDB
const CollectionName = "transports"

//TransportRepository is the MongoDB implementation of the interface domain.TransportRepository
type TransportRepository struct {
	Conn *mongo.Client
}

//Get a Transport by ID
func (repo TransportRepository) Get(id string) (*domain.Transport, error) {
	collection := repo.Conn.Database(DatabaseName).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{"_id": id}
	var transport = domain.Transport{}
	err := collection.FindOne(ctx, filter).Decode(&transport)
	if err != nil {
		return nil, err
	}
	return &transport, nil
}

//GetAll Transports
func (repo TransportRepository) GetAll() ([]*domain.Transport, error) {
	var transports []*domain.Transport
	transports = make([]*domain.Transport, 0)
	collection := repo.Conn.Database(DatabaseName).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result domain.Transport
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		transports = append(transports, &result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return transports, nil
}

//Save stores a Transport to the Database and returns the new Transport ID
func (repo TransportRepository) Save(transport *domain.Transport) (string, error) {
	collection := repo.Conn.Database(DatabaseName).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	transport.LastPositionReportedAt = time.Now()
	res, err := collection.InsertOne(ctx, transport)
	if err != nil {
		return "", err
	}
	id := res.InsertedID
	return id.(string), nil
}

//UpdatePosition updates the position of the transport with the given ID
func (repo TransportRepository) UpdatePosition(transportID string, newPosition *domain.Position) (*domain.Transport, error) {
	collection := repo.Conn.Database(DatabaseName).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{"_id": transportID}
	var transport = domain.Transport{}
	err := collection.FindOne(ctx, filter).Decode(&transport)
	if err != nil {
		return nil, err
	}
	transport.CurrentPosition = newPosition
	transport.LastPositionReportedAt = time.Now()
	update := bson.M{"$set": bson.M{"lastPositionReportedAt": transport.LastPositionReportedAt, "currentPosition": bson.M{"lat": newPosition.Lat, "lng": newPosition.Lng}}}
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	if res.ModifiedCount == 0 {
		return nil, errors.New("No documents updated")
	}

	return &transport, nil
}

func init() {
	repository.Repos.Add(repository.TransportRepositoryName, TransportRepository{Conn: DBClient.Conn})
}
