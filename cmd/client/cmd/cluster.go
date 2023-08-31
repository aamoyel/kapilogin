package cmd

import (
	"fmt"
	"log/slog"

	"github.com/spf13/cobra"

	"github.com/aamoyel/kapilogin/cmd/client/pkg/config"
)

var getClustersCmd = &cobra.Command{
	Use:   "clusters",
	Short: "clusters subcommand",
	Long:  `Get a list of Kubernetes clusters managed by Cluster API`,
	RunE:  runGet,
}

func init() {
	getCmd.AddCommand(getClustersCmd)
}

func runGet(cmd *cobra.Command, args []string) error {
	cfgUri, err := cmd.Flags().GetString("config")
	if err != nil {
		slog.Error("Error when getting 'config' flag value: %v", err)
	}
	config, err := config.GetConfig(cfgUri)
	if err != nil {
		slog.Error("Error at config retrieval: %w", err)
	}

	fmt.Printf("%+v\n", config)
	return nil
}
