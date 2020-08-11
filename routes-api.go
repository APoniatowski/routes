package main

import (
	"encoding/json"
	"net/http"
	"sort"
)

func routesEndpoint(w http.ResponseWriter, r *http.Request) {
	// check that routes is used in the url and send an error if it is malformed
	switch r.URL.Path {
	case "/routes":
		switch r.Method {
		case "GET":
			// if it matches routes, write 200 (OK) as header

			var routesResponse routesResponse
			// get src and dst queries
			routesResponse.Source = r.URL.Query().Get("src")
			destinationAdd := r.URL.Query()["dst"]

			// init (or make) the slice/array the length of the total number of destinations
			routesResponse.Routes = make([]Routes, len(destinationAdd))

			// iterate through the parsed query and add them to the struct
			for i := range destinationAdd {
				routesResponse.Routes[i].Destination = destinationAdd[i]
			}

			// call OSRM for distance and duration information
			routesResponse.routesAPICall()

			// sort distances and durations
			sort.SliceStable(routesResponse.Routes, func(i, j int) bool {
				return routesResponse.Routes[i].Distance < routesResponse.Routes[j].Distance
			})

			// marshal the struct to json
			jsonResponse, errMarshal := json.MarshalIndent(routesResponse, "", "  ")
			if errMarshal != nil {
				http.Error(w, errMarshal.Error(), http.StatusInternalServerError)
			}

			// send data to user
			w.Header().Set("Content-Type", "application/json")
			_, errResponse := w.Write(jsonResponse)
			if errResponse != nil {
				http.Error(w, errResponse.Error(), http.StatusInternalServerError)
			}
		default:
			// should only accept GET queries
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)

		}
	default:
		// if it does not match, send 404, not found
		http.NotFound(w, r)
	}
}
