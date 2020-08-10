package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *routesResponse) routesAPICall() {
	for i := range s.Routes {
		const (
			osrmURL      string = "http://router.project-osrm.org/route/v1/driving/"
			osrmOverview string = "?overview=false"
		)

		// var urlBuilder string
		urlBuilder := osrmURL + s.Source + ";" + s.Routes[i].Destination

		// call OSRM
		osrmDistance, osrmDuration := s.osrmCall(urlBuilder)

		// assign returned data to structs
		s.Routes[i].Duration = osrmDuration
		s.Routes[i].Distance = osrmDistance
	}
}

// call to OSRM to receive distance and duration data and return them
func (s *routesResponse) osrmCall(query string) (float64, float64) {
	// GET request
	response, errResponse := http.Get(query)
	if errResponse != nil {
		log.Println(errResponse.Error())
	}
	defer response.Body.Close()

	// unmarshal/decode the data
	var data map[string]interface{}
	errUnmarshal := json.NewDecoder(response.Body).Decode(&data)
	if errUnmarshal != nil {
		log.Println(errUnmarshal.Error())
	}

	// get data from routes
	addData := data["routes"]
	var osrmDistance float64
	var osrmDuration float64

	// assert/convert to float64, as the type is float64. not float32
	for _, val := range addData.([]interface{}) {
		osrmDistance = val.(map[string]interface{})["distance"].(float64)
		osrmDuration = val.(map[string]interface{})["duration"].(float64)
	}

	return osrmDistance, osrmDuration
}
