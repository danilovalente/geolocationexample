package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/danilovalente/geolocationexample/domain"
	"github.com/danilovalente/geolocationexample/usecase"
	"github.com/gorilla/mux"
)

//DefaultTransportPageSize defines the default number of Transports per page
const DefaultTransportPageSize = 10

func getTransport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var ID = vars["id"]
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	transport, err := usecase.GetTransport(ID)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(domain.GetErrorMessageBytes("Transports not found: ", err))
		return
	}

	transportJSON, err := json.Marshal(transport)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(domain.GetErrorMessageBytes("Marshaling error: ", err))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(transportJSON)
}

func getTransports(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	Transports, err := usecase.GetTransports()

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(domain.GetErrorMessageBytes("Transports not found: ", err))
		return
	}

	TransportsJSON, err := json.Marshal(Transports)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(domain.GetErrorMessageBytes("Marshaling error: ", err))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(TransportsJSON)
}

func postTransport(w http.ResponseWriter, r *http.Request) {
	var transport *domain.Transport
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(domain.GetErrorMessageBytes("An error occurred while trying to read the request body: ", err))
		return
	}
	err = json.Unmarshal(b, &transport)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(domain.GetErrorMessageBytes("Marshaling error: ", err))
		return
	}

	transportID, err := usecase.CreateTransport(transport)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(domain.GetErrorMessageBytes("An error occurred while trying to store the transport: ", err))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Location", "/geolocation/v1/transport/"+transportID)
}

func putPosition(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var transportID = vars["transportId"]
	var position *domain.Position
	var transport *domain.Transport
	w.Header().Set("Content-Type", "application/json")

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(domain.GetErrorMessageBytes("An error occurred while trying to read the request body: ", err))
		return
	}
	err = json.Unmarshal(b, &position)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(domain.GetErrorMessageBytes("Marshaling error: ", err))
		return
	}

	transport, err = usecase.UpdateTransportPosition(transportID, position)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(domain.GetErrorMessageBytes("An error occurred while trying to update the current position: ", err))
		return
	}
	TransportJSON, err := json.Marshal(transport)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(domain.GetErrorMessageBytes("Marshaling error: ", err))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(TransportJSON)
}
