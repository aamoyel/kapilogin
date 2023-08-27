package cluster

import (
	"context"
	"log"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"

	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

type Cluster struct {
	Name     string `json:"name"`
	Endpoint string `json:"endpoint"`
}

func GetClusterList() []Cluster {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	clusterGvk := schema.GroupVersionResource{Group: "cluster.x-k8s.io", Version: "v1beta1", Resource: "clusters"}
	cls, err := clientset.Resource(clusterGvk).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		log.Panicf("error getting resource from dynamic client, err: %v", err.Error())
	}

	var clusters clusterv1.ClusterList
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(cls.UnstructuredContent(), &clusters)
	if err != nil {
		log.Panicf("error converting unstructured obj to 'ClusterList' type, err: %v", err.Error())
	}

	var clusterList []Cluster
	for _, cluster := range clusters.Items {
		var c = Cluster{
			Name:     cluster.GetName(),
			Endpoint: cluster.Spec.ControlPlaneEndpoint.Host,
		}
		clusterList = append(clusterList, c)
	}

	return clusterList
}
