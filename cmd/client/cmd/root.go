package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kapilogin",
	Short: "List you available clusters and claim your kubeconfig files with ease",
	Long:  "kapilogin allows you to get and configure your kubeconfig files for available Kubernetes clusters managed by Cluster API",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var (
	cfgUri string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgUri, "config", "c", "", "Path to kapilogin configuration or to a remote url (i.e. https://example.com/kapilogin.yaml), KAPILOGIN_CONFIG environment variable overrides this flag")
}
