package main

import (
	"net/http"

	"k8s.io/klog"

	hu "hypercloud-multi-api-server/hypercloudurl"
	hcr "hypercloud-multi-api-server/hyperclusterresource"
	rc "hypercloud-multi-api-server/remotecluster"
)

func main() {
	// Req multiplexer
	mux := http.NewServeMux()
	mux.HandleFunc("/hyperclusterresource", hyperClusterResource)
	mux.HandleFunc("/hypercloudurl", hypercloudUrl)
	mux.HandleFunc("/remotecluster", remoteCluster)

	// HTTP Server Start
	klog.Info("Starting Hypercloud-Multi-Operator-API server...")
	if err := http.ListenAndServe(":80", mux); err != nil {
		klog.Errorf("Failed to listen and serve Hypercloud-Multi-Operator-API server: %s", err)
	}
	klog.Info("Started Hypercloud-Multi-Operator-API server")
}

func hypercloudUrl(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPut:
		hu.Put(res, req)
	default:
		//error
	}
}

func remoteCluster(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		rc.Get(res, req)
	default:
		//error
	}
}

func hyperClusterResource(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPut:
		hcr.Put(res, req)
	default:
		//error
	}
}
