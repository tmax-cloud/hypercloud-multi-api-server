package util

import (
	"flag"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog"
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
	var kubeconfig2 *string
	var err error
	var config *restclient.Config

	if home := homedir.HomeDir(); home != "" {
		kubeconfig2 = flag.String("kubeconfig2", filepath.Join(home, ".kube", "config"), "/root/.kube")
	} else {
		kubeconfig2 = flag.String("kubeconfig2", "", "/root/.kube")
	}
	flag.Parse()

	config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig2)
	if err != nil {
		klog.Errorln(err)
		panic(err)
	}
	config.Burst = 100
	config.QPS = 100

	if err != nil {
		klog.Errorln(err)
		panic(err)
	}

	setScheme()
	hostclient, err = client.New(config, client.Options{Scheme: scheme})

	// If api-server on POD, activate below code and delete above
	// creates the in-cluster config
	// var err error
	// config, err = restclient.InClusterConfig()
	// if err != nil {
	// 	panic(err.Error())
	// }
	// // creates the clientset
	// Clientset, err = kubernetes.NewForConfig(config)
	// if err != nil {
	// 	panic(err.Error())
	// }
}

func GetHostClient() *client.Client {
	return &hostclient
}
