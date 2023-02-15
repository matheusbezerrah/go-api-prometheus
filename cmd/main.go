package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	recordsCounterMetric()
	recordsGaugeMetric()
	recordsHistogramMetric()
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}

var (
	// COUNTER
	api_num_request_post = promauto.NewCounter(
		prometheus.CounterOpts{
			Name:        "api_request_total",
			Help:        "The total number of requests",
			ConstLabels: prometheus.Labels{"method": "POST", "handler": "/message"},
		},
	)
	api_num_request_get = promauto.NewCounter(
		prometheus.CounterOpts{
			Name:        "api_request_total",
			Help:        "The total number of requests",
			ConstLabels: prometheus.Labels{"method": "GET", "handler": "/message"},
		},
	)
	// GAUGE
	api_num_clients_online = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "api_clients_online",
			Help: "The number of clients online",
		},
	)
	// HISTOGRAM
	api_response_time = promauto.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "api_response_time_milliseconds",
			Help:    "Response time of the api in milliseconds",
			Buckets: prometheus.ExponentialBuckets(50, 2, 6),
		},
	)
	// SUMMARY
	api_summary_response_time = promauto.NewSummary(
		prometheus.SummaryOpts{
			Name:       "api_summary_response_time_milliseconds",
			Help:       "Response time of the api in milliseconds",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
	)
)

func recordsCounterMetric() {
	go func() {
		for {
			api_num_request_post.Inc()
			api_num_request_post.Add(10)
			time.Sleep(2 * time.Second)
			api_num_request_get.Inc()
			time.Sleep(1 * time.Second)
		}
	}()
}

func recordsGaugeMetric() {
	go func() {
		for {
			api_num_clients_online.Add(25)
			time.Sleep(10 * time.Second)
		}
	}()
	go func() {
		for {
			api_num_clients_online.Dec()
			time.Sleep(2 * time.Second)
		}
	}()
}

func recordsHistogramMetric() {

	go func() {
		for {
			start := time.Now()
			time.Sleep(45 * time.Millisecond)
			api_response_time.Observe(float64(time.Since(start).Milliseconds()))
			api_summary_response_time.Observe(float64(time.Since(start).Milliseconds()))
		}
	}()
	go func() {
		for {
			start := time.Now()
			time.Sleep(150 * time.Millisecond)
			api_response_time.Observe(float64(time.Since(start).Milliseconds()))
			api_summary_response_time.Observe(float64(time.Since(start).Milliseconds()))
		}
	}()
	go func() {
		for {
			start := time.Now()
			time.Sleep(210 * time.Millisecond)
			api_response_time.Observe(float64(time.Since(start).Milliseconds()))
			api_summary_response_time.Observe(float64(time.Since(start).Milliseconds()))
		}
	}()
	go func() {
		for {
			start := time.Now()
			time.Sleep(250 * time.Millisecond)
			api_response_time.Observe(float64(time.Since(start).Milliseconds()))
			api_summary_response_time.Observe(float64(time.Since(start).Milliseconds()))
		}
	}()
	go func() {
		for {
			start := time.Now()
			time.Sleep(300 * time.Millisecond)
			api_response_time.Observe(float64(time.Since(start).Milliseconds()))
			api_summary_response_time.Observe(float64(time.Since(start).Milliseconds()))
		}
	}()
	go func() {
		for {
			start := time.Now()
			time.Sleep(370 * time.Millisecond)
			api_response_time.Observe(float64(time.Since(start).Milliseconds()))
			api_summary_response_time.Observe(float64(time.Since(start).Milliseconds()))
		}
	}()
	go func() {
		for {
			start := time.Now()
			time.Sleep(405 * time.Millisecond)
			api_response_time.Observe(float64(time.Since(start).Milliseconds()))
			api_summary_response_time.Observe(float64(time.Since(start).Milliseconds()))
		}
	}()
	go func() {
		for {
			start := time.Now()
			time.Sleep(748 * time.Millisecond)
			api_response_time.Observe(float64(time.Since(start).Milliseconds()))
			api_summary_response_time.Observe(float64(time.Since(start).Milliseconds()))
		}
	}()
}
