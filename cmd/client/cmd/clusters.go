package cmd

import (
	"fmt"
	"log/slog"

	"github.com/spf13/cobra"

	cl "github.com/aamoyel/kapilogin/cmd/client/pkg/clusterslist"
	"github.com/aamoyel/kapilogin/cmd/client/pkg/config"
)

var getClustersCmd = &cobra.Command{
	Use:   "clusters",
	Short: "Get a list of Kubernetes clusters managed by Cluster API",
	RunE:  runClusters,
}

func init() {
	getCmd.AddCommand(getClustersCmd)
}

func runClusters(cmd *cobra.Command, args []string) error {
	cfgUri, err := cmd.Flags().GetString("config")
	if err != nil {
		return err
	}
	config, err := config.GetConfig(cfgUri)
	if err != nil {
		slog.Error("when getting configuration", err)
	}

	clustersList, err := cl.GetClusters(config)
	if err != nil {
		slog.Error("when getting clusters list", err)
	}

	fmt.Printf("%-24s %-10s\n", "NAME", "ENDPOINT")
	for _, cluster := range clustersList {
		fmt.Printf("%-24s %-10s\n", cluster.Name, cluster.Endpoint)
	}

	return nil
}
