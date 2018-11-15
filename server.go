package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	histogram := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "hash_seconds",
		Help: "Time taken to create hashes",
	}, []string{"code"})

	r := mux.NewRouter()

	r.Handle("/metrics", prometheusHandler())
	r.Handle("/hash", hashHandler(histogram))

	prometheus.Register(histogram)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func prometheusHandler() http.Handler {
	return promhttp.Handler()
}

func hashHandler(histogram *prometheus.HistogramVec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer r.Body.Close()
		code := http.StatusInternalServerError

		defer func() {
			duration := time.Since(start)
			histogram.WithLabelValues(fmt.Sprintf("%d", code)).Observe(duration.Seconds())
		}()

		code = http.StatusMethodNotAllowed
		if r.Method != http.MethodPost {
			w.WriteHeader(code)

			return
		}

		code = http.StatusOK
		w.WriteHeader(code)
		body, _ := ioutil.ReadAll(r.Body)

		fmt.Printf("\"%s\"\n", string(body))

		hashed := computeSum(body)
		w.Write(hashed)

	}
}

func computeSum(body []byte) []byte {
	h := sha256.New()

	h.Write(body)
	hashed := hex.EncodeToString(h.Sum(nil))
	return []byte(hashed)
}
