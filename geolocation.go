package main

import (
	"log"
	"net/http"

	"github.com/danilovalente/geolocationexample/config"
	"github.com/danilovalente/geolocationexample/controller"
)

func main() {
	router, err := controller.MapHandlers()
	if err != nil {
		panic(err)
	}

	log.Println("geolocation-api started listening in :" + config.Port)
	http.ListenAndServe(":"+config.Port, router)

}
