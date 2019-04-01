package mongodb

import (
	"context"
	"log"

	"github.com/danilovalente/geolocationexample/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoClient is a struct to keep the DB connection and URI
type MongoClient struct {
	DBURI string
	Conn  *mongo.Client
}

//DBClient holds a client connection with the database
var DBClient MongoClient

func init() {
	DBClient = MongoClient{}
	DBClient.DBURI = config.DBConnectionString
	client, err := mongo.NewClient(options.Client().ApplyURI(DBClient.DBURI))
	if err != nil {
		log.Fatal("An error occurred while trying to open a DB connection")
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("An error occurred while trying to open a DB connection")
	}
	DBClient.Conn = client
	log.Println("Connected to database " + DBClient.DBURI)
}
