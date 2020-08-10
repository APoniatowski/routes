package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/routes", routesEndpoint)

	err := http.ListenAndServe(envPort(), nil)
	if err != nil {
		fmt.Println(" Error starting HTTP server...")
		fmt.Println("<=============================>")
		log.Println(err.Error())
	}
}
