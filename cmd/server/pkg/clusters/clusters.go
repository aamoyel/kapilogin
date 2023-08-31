package cluster

import (
	"context"
	"log/slog"
	"strconv"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

type Cluster struct {
	Name     string `json:"name"`
	Endpoint string `json:"endpoint"`
	CaCert   string `json:"cacert"`
}

func GetClusterList() []Cluster {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// creates the dyn client
	dynclient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	clusterGvk := schema.GroupVersionResource{Group: "cluster.x-k8s.io", Version: "v1beta1", Resource: "clusters"}
	cls, err := dynclient.Resource(clusterGvk).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		slog.Error("error getting resource from dynamic client, err: %+v", err.Error())
	}

	var clusters clusterv1.ClusterList
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(cls.UnstructuredContent(), &clusters)
	if err != nil {
		slog.Error("error converting unstructured obj to 'ClusterList' type, err: %+v", err.Error())
	}

	var clusterList []Cluster
	for _, cluster := range clusters.Items {
		secret, err := clientset.CoreV1().Secrets(cluster.GetNamespace()).Get(context.TODO(), cluster.GetName()+"-ca", metav1.GetOptions{})
		if err != nil {
			slog.Error("Error when getting ca cert secret: %+v", err)
		}

		var c = Cluster{
			Name:     cluster.GetName(),
			Endpoint: "https://" + cluster.Spec.ControlPlaneEndpoint.Host + ":" + strconv.Itoa(int(cluster.Spec.ControlPlaneEndpoint.Port)),
			CaCert:   string(secret.Data["tls.crt"]),
		}
		clusterList = append(clusterList, c)
	}

	return clusterList
}
