package main

import (
	"fmt"
	"net/http"
	"newsfeeder/httpd/handler"
	"newsfeeder/repository/newsfeeds"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Define a counter metric
var requestsProcessed = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "myapp_requests_processed_total",
		Help: "The total number of requests processed.",
	})

// Define a histogram metric
var requestDuration = prometheus.NewHistogram(
	prometheus.HistogramOpts{
		Name: "myapp_request_duration_seconds",
		Help: "The duration of requests in seconds.",
	})

func init() {
	prometheus.MustRegister(requestsProcessed, requestDuration)
}
func main() {
	feed := newsfeeds.New()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		requestsProcessed.Inc()
		start := time.Now()

		handler.PingGet()(c) // Pass feed as an argument to GetNews
		duration := time.Since(start).Seconds()
		requestDuration.Observe(duration)
	})
        

	r.GET("/news", func(c *gin.Context) {
		requestsProcessed.Inc()
		start := time.Now()

		handler.GetNews(feed)(c) // Pass feed as an argument to GetNews
		duration := time.Since(start).Seconds()
		requestDuration.Observe(duration)
	})

	r.POST("/addNews", func(c *gin.Context) {
		requestsProcessed.Inc()
		start := time.Now()

		handler.PostNews(feed)(c) // Pass feed as an argument to PostNews
		duration := time.Since(start).Seconds()
		requestDuration.Observe(duration)
	})

	r.GET("/metric", gin.WrapH(promhttp.Handler())) // Expose metrics endpoint
	r.GET("/health", func(c *gin.Context) {
		requestsProcessed.Inc()
		start := time.Now()

		healthHandler(c.Writer,c.Request) // Pass feed as an argument to PostNews
		duration := time.Since(start).Seconds()
		requestDuration.Observe(duration)
	})

	r.Run(":8080")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Only GET method is allowed")
		return
	}

	// Perform health checks based on metrics
	if isHealthy() {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Application is healthy")
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprint(w, "Application is unhealthy")
	}
}

func isHealthy() bool {
	// Implement your health check logic here
	// Evaluate the collected metrics and compare them against thresholds
	// Return true if the application is healthy, false otherwise

	// Example: Check if the error rate is below a certain threshold
	errorRate := getErrorRateMetric()
	if errorRate <= 0.05 { // Assuming a threshold of 5%
		return true
	}
	return false
}

func getErrorRateMetric() float64 {
	// Retrieve and calculate the error rate metric
	// Example: Calculate the error rate based on the total requests and error count
	totalRequests := getTotalRequestsMetric()
	errorCount := getErrorCountMetric()
	if totalRequests > 0 {
		return float64(errorCount) / float64(totalRequests)
	}
	return 0.0
}

func getTotalRequestsMetric() int {
	// Retrieve the total requests metric
	// Example: Get the value from a metrics storage or monitoring system
	return 1000
}

func getErrorCountMetric() int {
	// Retrieve the error count metric
	// Example: Get the value from a metrics storage or monitoring system
	return 50
}
