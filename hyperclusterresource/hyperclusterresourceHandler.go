package hyperclusterresource

import (
	"context"
	"encoding/json"
	"hypercloud-multi-api-server/util"
	"io/ioutil"
	"net/http"
	"strings"

	"k8s.io/klog"

	hyperv1 "hypercloud-multi-api-server/external/hyper/v1"

	corev1 "k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/cluster-api/util/patch"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

/*
  Input
  - query param: cluster-name
  - body:
    {
	    "resourcelist" : [
	        {
	          "name":"pod",
	          "Total": 1,
	          "Running": 1
	        },
	      	{
	          "name":"cpu",
	          "Total": 1,
	          "Running": 1
	        },
	      	{
	          "name":"memory",
	          "Total": 1,
	          "Running": 1
	        },
	      	{
	          "name":"storage",
	          "Total": 1,
	          "Running": 1
	        }
	    ]
	}
*/
func Put(res http.ResponseWriter, req *http.Request) {
	klog.Infoln("**** PUT/hyperclusterresource")

	var resourceList hyperv1.ResourceTypeList

	queryParams := req.URL.Query()
	clusterName := queryParams.Get(util.HYPERCLUSTERRESOURCE_CLUSTER_NAME)
	payloadBytes, _ := ioutil.ReadAll(req.Body)
	if err := json.Unmarshal(payloadBytes, &resourceList); err != nil {
		klog.Errorln(err)
	}

	//get hostConfig
	hostclient := util.GetHostClient()

	//update cpu, memory, storage, pod usage
	updateHCR(*hostclient, clusterName, resourceList.ResourceList)
}

func updateHCR(c client.Client, clusterName string, resourceList []hyperv1.ResourceType) {
	hcr := &hyperv1.HyperClusterResource{}

	if err := c.Get(context.TODO(), types.NamespacedName{Name: clusterName, Namespace: util.HYPERCLUSTERRESOURCE_NAMESPACE}, hcr); err != nil {
		klog.Errorln(err)
	}

	//set helper
	helper, _ := patch.NewHelper(hcr, c)
	defer func() {
		helper.Patch(context.TODO(), hcr)
	}()

	//init hcr.Status.Resources from configmap in kube-federation-system namespace
	if hcr.Status.Resources == nil {
		hcr.Status.Resources = []hyperv1.ResourceType{}
		initHCR(c, hcr)
	}

	//handling allocate resource info
	for _, resource := range resourceList {
		for index, value := range hcr.Status.Resources {
			if strings.Compare(strings.ToLower(value.Name), strings.ToLower(resource.Name)) == 0 {
				hcr.Status.Resources[index].Total = value.Total + resource.Total
				hcr.Status.Resources[index].Running = value.Running + resource.Running
			}
		}
	}
}

func initHCR(c client.Client, hcr *hyperv1.HyperClusterResource) {
	configmap := &corev1.ConfigMap{}

	if err := c.Get(context.TODO(), types.NamespacedName{Name: util.HYPERCLUSTERRESOURCE_CONFIGMAP_NAME, Namespace: util.HYPERCLUSTERRESOURCE_NAMESPACE}, configmap); err != nil {
		klog.Errorln(err)
	}

	for key, _ := range configmap.Data {
		hcr.Status.Resources = append(hcr.Status.Resources, hyperv1.ResourceType{Name: key, Total: 0, Running: 0})
	}
}
