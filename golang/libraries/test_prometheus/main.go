package main

import (
	"flag"
	"log"
	"net/http"
	"time"
	"math"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var addr = flag.String("listen-address", ":58080", "The address to listen on for HTTP requests.")

func main() {
	go reportCounter()
	go reportGauge()
	go reportHistogram()
	go reportSummary()

	flag.Parse()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func reportCounter() {
	test_counter := promauto.NewCounter(prometheus.CounterOpts{
		Namespace:"my_namespace",
		Subsystem:"my_subsystem",
		Name:"my_test_counter",
		Help:"my_help",
		ConstLabels:prometheus.Labels{"project":"test_prometheus"},
	})

	for {
		test_counter.Inc()

		time.Sleep(time.Millisecond * 100)
	}
}

func reportGauge() {
	test_gauge := promauto.NewGauge(prometheus.GaugeOpts{
		Namespace:"my_namespace",
		Subsystem:"my_subsystem",
		Name:"my_test_gauge",
		Help:"my_help",
		ConstLabels:prometheus.Labels{"project":"test_prometheus"},
	})

	for {
		test_gauge.Set(50.0)

		time.Sleep(time.Millisecond * 100)
	}
}

func reportHistogram() {
	test_histogram := promauto.NewHistogram(prometheus.HistogramOpts{
		Namespace:"my_namespace",
		Subsystem:"my_subsystem",
		Name:"my_test_histogram",
		Help:"my_help",
		ConstLabels:prometheus.Labels{"project":"test_prometheus"},
		Buckets:prometheus.LinearBuckets(20, 5, 5),
	})

	index := 0

	for {
		if index > 1000 {
			index = 0
		}

		index++

		test_histogram.Observe(30 + math.Floor(120*math.Sin(float64(index)*0.1))/10)

		time.Sleep(time.Millisecond * 100)
	}
}

func reportSummary() {
	test_summary := promauto.NewSummary(prometheus.SummaryOpts{
		Namespace:"my_namespace",
		Subsystem:"my_subsystem",
		Name:"my_test_summary",
		Help:"my_help",
		ConstLabels:prometheus.Labels{"project":"test_prometheus"},
		Objectives:map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	})

	index := 0

	for {
		if index > 1000 {
			index = 0
		}

		index++

		test_summary.Observe(30 + math.Floor(120*math.Sin(float64(index)*0.1))/10)

		time.Sleep(time.Millisecond * 100)
	}
}
