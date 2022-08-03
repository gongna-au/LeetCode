package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func main() {
	opsQueued := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "our_company",
			Subsystem: "blob_storage",
			Name:      "ops_queued",
			Help:      "Number of blob storage operations waiting to be processed, partitioned by user and type.",
		},
		[]string{
			// Which user has requested the operation?
			"user",
			// Of what type is the operation?
			"type",
		},
	)

	prometheus.MustRegister(opsQueued)

	// Increase a value using compact (but order-sensitive!) WithLabelValues().
	opsQueued.WithLabelValues("bob", "put").Add(4)
	opsQueued.WithLabelValues("tom", "put").Add(7)
	// Increase a value with a map using WithLabels. More verbose, but order
	// doesn't matter anymore.
	opsQueued.With(prometheus.Labels{"type": "delete", "user": "alice"}).Inc()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
