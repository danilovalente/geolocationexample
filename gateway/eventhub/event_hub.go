package eventhub

import (
	"context"
	"log"

	"github.com/Azure/azure-amqp-common-go/aad"
	eventhubs "github.com/Azure/azure-event-hubs-go"
	"github.com/danilovalente/geolocationexample/config"
)

/*EventHub is an struct that holds an Event Hub configuration
and have the features of sending and consuming events*/
type EventHub struct {
	namespace string
	name      string
	hub       *eventhubs.Hub
}

var eventHub *EventHub

//Send a new message to the configured Event Hub
func (eventHub *EventHub) Send(message string) error {
	return eventHub.hub.Send(context.Background(), eventhubs.NewEventFromString(message))
}

func init() {
	eventHub = &EventHub{namespace: config.EventHubNamespace, name: config.EventHubName}
	tokenProvider, err := aad.NewJWTProvider(aad.JWTProviderWithEnvironmentVars())
	if err != nil {
		log.Fatalf("failed to configure AAD JWT provider: %s\n", err)
	}
	eventHub.hub, err = eventhubs.NewHub(config.EventHubNamespace, config.EventHubName, tokenProvider)
	if err != nil {
		log.Fatalf("failed to get hub %s\n", err)
	}
}
