package config

import (
	"os"
)

var (
	// DBConnectionString holds the database Connection String
	DBConnectionString string
	// HashVersion keeps the MD5 hash version
	HashVersion string
	//EventHubNamespace holds the EventHub namespace
	EventHubNamespace string
	//EventHubName holds the name of the EventHub
	EventHubName string
	//Port contains the port in which the application listens
	Port string
)

func init() {
	DBConnectionString = os.Getenv("DBCONNECTIONSTRING")
	HashVersion = os.Getenv("HASHVERSION")
	EventHubNamespace = os.Getenv("EVENTHUBNAMESPACE")
	EventHubName = os.Getenv("EVENTHUBNAME")
	Port = os.Getenv("PORT")
}
