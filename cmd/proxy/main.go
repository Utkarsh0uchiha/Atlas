package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	target, _ := url.Parse("http://localhost:8081")

	proxy := httputil.NewSingleHostReverseProxy(target)

	http.ListenAndServe(":8080", proxy)
}
