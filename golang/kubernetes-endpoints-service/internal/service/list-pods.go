package service

import (
	"github.com/parthw/kubernetes-endpoints-service/internal/logger"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// CountPods to count number of pods
func CountPods(inCluster bool) {
	var clientset *kubernetes.Clientset
	clientset = newKubernetesClient(clientset, inCluster)
	if clientset == nil {
		return
	}

	pods, err := clientset.CoreV1().Pods("default").List(metav1.ListOptions{})
	if err != nil {
		logger.Log.Errorf("Failed to fetch pods - ", err.Error())
	}
	logger.Log.Infof("There are %d pods in the cluster\n", len(pods.Items))
}
