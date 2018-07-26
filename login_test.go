package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginBasicAuth(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set("www-authenticate", `Basic realm="Registry Realm",service="Docker registry"`)
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}))
	defer ts.Close()

	run(t, "login", "-u", "j3ss", "-p", "ss3j", "-insecure-registry", ts.URL)
}
