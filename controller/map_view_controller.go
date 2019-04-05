package controller

import (
	"io/ioutil"
	"net/http"
)

func loadFile(path string) ([]byte, error) {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func getMapView(w http.ResponseWriter, r *http.Request) {
	var path = ""
	path = "MapView.html"
	p, _ := loadFile(path)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")
	w.Write(p)
}
