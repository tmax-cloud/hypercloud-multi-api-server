package remotecluster

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"hypercloud-multi-api-server/util"
)

var (
	scheme       = runtime.NewScheme()
	remoteClient client.Client
)

/*
  Input
  - query param: cluster-name, kind
  Output
  - body: pod list
*/
func Get(res http.ResponseWriter, req *http.Request) {
	klog.Infoln("**** GET/remotecluster")

	queryParams := req.URL.Query()
	clusterName := queryParams.Get(util.HYPERCLUSTERRESOURCE_CLUSTER_NAME)
	kind := strings.ToLower(queryParams.Get(util.REMOTERESOURCE_KIND))

	//get remoteConfig frome secret
	hostclient := util.GetHostClient()
	if restConfig, err := getConfigFromSecret(*hostclient, clusterName); err != nil {
		klog.Error(err)
		return
	} else {
		utilruntime.Must(corev1.AddToScheme(scheme))
		remoteClient, err = client.New(restConfig, client.Options{Scheme: scheme})
	}

	//get resource frome remote cluster and make response body
	enc := json.NewEncoder(res)
	res.Header().Set("Content-Type", "application/json")
	switch kind {
	case util.REMOTERESOURCE_KIND_POD:
		enc.Encode(getPodList(remoteClient))
	}
}

func getPodList(r client.Client) *corev1.PodList {
	podList := &corev1.PodList{}
	r.List(context.TODO(), podList)

	return podList
}

func getConfigFromSecret(c client.Client, clusterName string) (*restclient.Config, error) {
	secret := &corev1.Secret{}

	if err := c.Get(context.TODO(), types.NamespacedName{Name: clusterName + util.KUBECONFIG_POSTFIX, Namespace: util.KUBECONFIG_NAMESPACE}, secret); err != nil {
		klog.Errorln(err)
	}

	return getKubeConfig(*secret)
}

func getKubeConfig(s corev1.Secret) (*restclient.Config, error) {
	if value, ok := s.Data["value"]; ok {
		if clientConfig, err := clientcmd.NewClientConfigFromBytes(value); err == nil {
			if restConfig, err := clientConfig.ClientConfig(); err == nil {
				return restConfig, nil
			}
		}
	}
	return nil, errors.NewBadRequest("getClientConfig Error")
}
