package main

import (
	"bridge/bridge/libs/k8s"
	"bridge/bridge/log"
	"bridge/bridge/views/router"
	"net/http"
)

func main() {

	err := k8s.InitK8SClient()
	if err != nil {
		log.L.Fatalln("init k8s client failed", err)
	}

	r := router.Load()
	http.ListenAndServe("localhost:8080", r)
}
