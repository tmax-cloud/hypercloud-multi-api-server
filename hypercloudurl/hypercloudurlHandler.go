package hypercloudurl

import (
	"context"
	"hypercloud-multi-api-server/util"
	"net/http"

	"k8s.io/klog"

	hyperv1 "hypercloud-multi-api-server/external/hyper/v1"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/cluster-api/util/patch"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func Put(res http.ResponseWriter, req *http.Request) {
	klog.Infoln("**** PUT/hypercloudurl")

	//get queryParams
	queryParams := req.URL.Query()
	clusterName := queryParams.Get(util.HYPERCLUSTERRESOURCE_CLUSTER_NAME)
	clusterUrl := queryParams.Get(util.HYPERCLOUD_URL)

	//get hostConfig
	hostclient := util.GetHostClient()

	//update hypercluster url
	updateHCR(*hostclient, clusterName, clusterUrl)
}

func updateHCR(c client.Client, clusterName string, clusterUrl string) {
	hcr := &hyperv1.HyperClusterResource{}

	if err := c.Get(context.TODO(), types.NamespacedName{Name: clusterName, Namespace: util.HYPERCLUSTERRESOURCE_NAMESPACE}, hcr); err != nil {
		klog.Errorln(err)
	}

	//set helper
	helper, _ := patch.NewHelper(hcr, c)
	defer func() {
		helper.Patch(context.TODO(), hcr)
	}()

	if hcr.Annotations == nil {
		hcr.Annotations = map[string]string{}
	}
	hcr.Annotations[util.HYPERCLOUD_URL] = clusterUrl
}
