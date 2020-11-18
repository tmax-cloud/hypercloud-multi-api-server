package util

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	restclient "k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"

	hyperv1 "hypercloud-multi-api-server/external/hyper/v1"
)

var (
	scheme     = runtime.NewScheme()
	hostclient client.Client
)

func setScheme() {
	utilruntime.Must(hyperv1.AddToScheme(scheme))

	utilruntime.Must(corev1.AddToScheme(scheme))
}

func init() {
	// If api-server on POD, activate below code and delete above
	// creates the in-cluster config
	config, err := restclient.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	setScheme()
	hostclient, err = client.New(config, client.Options{Scheme: scheme})
}

func GetHostClient() *client.Client {
	return &hostclient
}
