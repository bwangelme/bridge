package main

import (
	"bridge/bridge/views/router"
	"net/http"
)

func main() {
	r := router.Load()
	http.ListenAndServe("localhost:8080", r)
}
