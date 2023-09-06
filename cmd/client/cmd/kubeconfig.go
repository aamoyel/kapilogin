package cmd

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/aamoyel/kapilogin/cmd/client/pkg/config"
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	cl "github.com/aamoyel/kapilogin/cmd/client/pkg/clusterslist"
)

var getKubeconfigCmd = &cobra.Command{
	Use:   "kubeconfig",
	Short: "kubeconfig subcommand",
	Long:  `Get the Kubeconfig of a Kubernetes cluster managed by Cluster API`,
	RunE:  runKubeconfig,
}

var (
	clusterName string
	clientCfg   clientcmdapi.Config
	kcCluster   clientcmdapi.Cluster
	kcContext   clientcmdapi.Context
	kcUser      clientcmdapi.AuthInfo
)

func init() {
	getCmd.AddCommand(getKubeconfigCmd)
	getKubeconfigCmd.Flags().StringVarP(&clusterName, "name", "n", "", "Name of the cluster")
	getKubeconfigCmd.MarkFlagRequired("name")
}

func runKubeconfig(cmd *cobra.Command, args []string) error {
	cfgUri, err := cmd.Flags().GetString("config")
	if err != nil {
		return err
	}
	config, err := config.GetConfig(cfgUri)
	if err != nil {
		slog.Error("when getting configuration", err)
	}

	clusterName, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}

	kubeconfig, err := getKubeconfig(config, clusterName)
	if err != nil {
		slog.Error("when getting kubeconfig", err)
	}

	fmt.Printf("%+v", kubeconfig)
	return nil
}

func getKubeconfig(config *config.Config, clusterName string) (string, error) {
	clustersList, err := cl.GetClusters(config)
	if err != nil {
		return "", err
	}

	for _, cluster := range clustersList {
		if cluster.Name == clusterName {
			clientCfg.APIVersion = "v1"
			clientCfg.Kind = "Config"

			kcCluster.Server = cluster.Endpoint
			kcCluster.CertificateAuthorityData = []byte(cluster.CaCert)
			clientCfg.Clusters = map[string]*clientcmdapi.Cluster{clusterName: &kcCluster}

			kcContext.Cluster = clusterName
			kcContext.AuthInfo = "oidc"
			clientCfg.Contexts = map[string]*clientcmdapi.Context{"oidc@" + clusterName: &kcContext}

			kcUser.Exec = &clientcmdapi.ExecConfig{
				APIVersion: "client.authentication.k8s.io/v1beta1",
				Command:    "kubectl",
			}
			execArgs := []string{
				"oidc-login",
				"get-token",
				"--oidc-issuer-url=" + config.OidcIssuerUrl,
				"--oidc-client-id=" + config.OidcClientId,
			}
			if config.OidcClientSecret != "" {
				execArgs = append(execArgs, "--oidc-client-secret="+config.OidcClientSecret)
			}
			kcUser.Exec.Args = execArgs
			clientCfg.AuthInfos = map[string]*clientcmdapi.AuthInfo{"oidc": &kcUser}

			kubeconfig, err := clientcmd.Write(clientCfg)
			if err != nil {
				return "", err
			}
			return string(kubeconfig), nil
		}
	}
	return "", errors.New("your cluster doesn't exist !")
}
