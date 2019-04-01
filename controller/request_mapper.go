package controller

import "github.com/gorilla/mux"

// MapHandlers function does the mapping from the request URLs to the correspondent handler functions
func MapHandlers() (*mux.Router, error) {
	router := mux.NewRouter()
	context := router.PathPrefix("/geolocation/v1").Subrouter()
	context.Methods("OPTIONS").HandlerFunc(corsHandler)
	context.Methods("GET").Path("/health").HandlerFunc(checkHealth)
	context.Methods("PUT").Path("/transport/{transportId}/position").HandlerFunc(putPosition)
	context.Methods("GET").Path("/transport/{id}").HandlerFunc(getTransport)
	context.Methods("GET").Path("/transport").HandlerFunc(getTransports)
	context.Methods("POST").Path("/transport").HandlerFunc(postTransport)
	return router, nil
}
