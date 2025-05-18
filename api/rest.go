package api

import (
	"fmt"
	"net/http"
)

func RegisterRest(mux *http.ServeMux) {
	mux.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from REST")
	})
}
