package clients

import (
	"os"

	chaosClient "github.com/litmuschaos/chaos-operator/pkg/client/clientset/versioned/typed/litmuschaos/v1alpha1"
	"github.com/pkg/errors"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// ClientSets is a collection of clientSets and kubeConfig needed
type ClientSets struct {
	KubeClient    *kubernetes.Clientset
	LitmusClient  *chaosClient.LitmuschaosV1alpha1Client
	KubeConfig    *rest.Config
	DynamicClient dynamic.Interface
}

// GenerateClientSetFromKubeConfig will generation both ClientSets (k8s, and Litmus) as well as the KubeConfig
func (clientSets *ClientSets) GenerateClientSetFromKubeConfig() error {

	config, err := getKubeConfig()
	if err != nil {
		return err
	}
	litmusClientSet, err := generateLitmusClientSet(config)
	if err != nil {
		return err
	}

	clientSets.LitmusClient = litmusClientSet
	clientSets.KubeConfig = config
	return nil
}

// getKubeConfig setup the config for access cluster resource
func getKubeConfig() (*rest.Config, error) {

	KubeConfig := os.Getenv("KUBECONFIG")
	// Use in-cluster config if kubeconfig path is not specified
	homeDir, err := os.UserHomeDir()
	if fileExists(homeDir + "/.kube/config") {
		KubeConfig = homeDir + "/.kube/config"
	}
	if KubeConfig == "" {
		return rest.InClusterConfig()
	}
	config, err := clientcmd.BuildConfigFromFlags("", KubeConfig)
	if err != nil {
		return config, err
	}
	return config, err
}

// generateK8sClientSet will generation k8s client
func generateK8sClientSet(config *rest.Config) (*kubernetes.Clientset, error) {
	k8sClientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to generate kubernetes clientSet, err: %v: ", err)
	}
	return k8sClientSet, nil
}

// generateLitmusClientSet will generate a LitmusClient
func generateLitmusClientSet(config *rest.Config) (*chaosClient.LitmuschaosV1alpha1Client, error) {
	litmusClientSet, err := chaosClient.NewForConfig(config)
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to create LitmusClientSet, err: %v", err)
	}
	return litmusClientSet, nil
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
