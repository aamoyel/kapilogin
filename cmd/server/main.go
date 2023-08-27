package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aamoyel/kapilogin/pkg/api"
)

func main() {
	// Configure application port if env var is set
	appPort, exists := os.LookupEnv("APP_PORT")
	if exists {
		appPort = ":" + appPort
	} else {
		appPort = ":8080"
	}

	http.HandleFunc("/", api.ClusterHandler)
	log.Fatal(http.ListenAndServe(appPort, nil))
}
