package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"metrics-export/export"
	"net/http"
	"os"
)

var (
	// Since we are dealing with custom Collector implementations, it might
	// be a good idea to try it out with a pedantic registry
	reg = prometheus.NewPedanticRegistry()
)

func register() {
	workerDB := export.NewClusterManager("cb")
	workerCA := export.NewClusterManager("ca")
	reg.MustRegister(workerDB)
	reg.MustRegister(workerCA)
}

func main() {
	register()
	gatherers := prometheus.Gatherers{
		prometheus.DefaultGatherer,
		reg,
	}

	h := promhttp.HandlerFor(gatherers,
		promhttp.HandlerOpts{
			ErrorLog:      log.Default(),
			ErrorHandling: promhttp.ContinueOnError,
		})
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("Error occur when start server %v", err)
		os.Exit(1)
	}
}
