package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Handle("/metrics", prometheusHandler())
	r.Handle("/hash", hashHandler())

	log.Fatal(http.ListenAndServe(":8002", r))
}

func prometheusHandler() http.Handler {
	return prometheus.Handler()
}

func hashHandler() {

}
