package service

import (
	"path/filepath"

	"github.com/parthw/kubernetes-endpoints-service/internal/logger"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func newKubernetesClient(clientset *kubernetes.Clientset, inCluster bool) *kubernetes.Clientset {
	if inCluster {
		clientset, _ = newInClusterClient()
	} else {
		clientset, _ = newOutClusterClient()
	}
	return clientset
}

func newInClusterClient() (*kubernetes.Clientset, error) {
	// rest.InClusterConfig uses service account token mounted at
	// /var/run/secrets/kubernetes.io/serviceaccount
	config, err := rest.InClusterConfig()
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}
	return clientset, nil
}

func newOutClusterClient() (*kubernetes.Clientset, error) {
	home := homedir.HomeDir()
	kubeconfig := filepath.Join(home, ".kube", "config")

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}
	return clientset, nil
}
