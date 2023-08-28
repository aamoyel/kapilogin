package cmd

import (
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

func runGet(cmd *cobra.Command, args []string) (err error) {
	cfgUri, _ = cmd.Flags().GetString("config")
	_, err = config.GetConfig(cfgUri)
	return nil
}
