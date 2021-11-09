package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"metrics-export/export"
	"net/http"
	"os"
)

var (
	// Since we are dealing with custom Collector implementations, it might
	// be a good idea to try it out with a pedantic registry
	reg = prometheus.NewPedanticRegistry()
)

func init() {
	workerDB := export.NewClusterManager("cb")
	workerCA := export.NewClusterManager("ca")
	reg.MustRegister(workerDB)
	reg.MustRegister(workerCA)

	// gauge
	reg.MustRegister(export.TempGauge)
	export.TempGauge.Set(24)
}

func main() {
	gatherers := prometheus.Gatherers{reg}
	h := promhttp.HandlerFor(
		gatherers,
		promhttp.HandlerOpts{
			ErrorLog:      logrus.New(),
			ErrorHandling: promhttp.ContinueOnError,
		})
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
	if err := http.ListenAndServe(":80", nil); err != nil {
		logrus.Println("Error occur when start server %v", err)
		os.Exit(1)
	}
}
