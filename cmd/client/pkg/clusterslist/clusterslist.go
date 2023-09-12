package clusterslist

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/aamoyel/kapilogin/cmd/client/pkg/config"
	cluster "github.com/aamoyel/kapilogin/cmd/server/pkg/clusters"
)

func GetClusters(config *config.Config) ([]*cluster.Cluster, error) {
	response, err := http.Get(config.KapiloginApiEndpoint)
	if err != nil {
		return nil, err
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	var responseObject []*cluster.Cluster
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return nil, err
	}

	return responseObject, nil
}
