package main

import (
	"os"
)

/*
	get the environment var for the port you'd prefer to use via ENV in k8s/dockerfile,
	or alternatively it will default to 80, if no port was specified
*/
func envPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "80"
	}

	return ":" + port
}
