package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aamoyel/kapilogin/cmd/server/pkg/cluster"
)

func main() {
	// Configure application port if env var is set
	appPort, exists := os.LookupEnv("APP_PORT")
	if exists {
		appPort = ":" + appPort
	} else {
		appPort = ":8080"
	}

	http.HandleFunc("/", ClusterHandler)
	log.Fatal(http.ListenAndServe(appPort, nil))
}

func ClusterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Just send out the JSON version of 'tom'
		j, _ := json.Marshal(cluster.GetClusterList())
		w.Write(j)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Printf("Method not allowed, Request: %v", r)
	}
}
