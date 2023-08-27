package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aamoyel/kapilogin/pkg/cluster"
)

func ClusterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Just send out the JSON version of 'tom'
		j, _ := json.Marshal(cluster.GetClusterList())
		w.Write(j)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Printf("Method not allowed, Request: %v",r)
	}
}
