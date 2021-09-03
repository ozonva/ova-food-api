package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	createCounter *prometheus.CounterVec
	updateCounter *prometheus.CounterVec
	removeCounter *prometheus.CounterVec
)

func RegisterMetrics() {

	createCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "createCounter"},
		[]string{"operation"})

	updateCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "updateCounter"},
		[]string{"operation"})

	removeCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "removeCounter"},
		[]string{"operation"})
}

func CounterIncrement(operation string) {
	if createCounter != nil {
		switch operation {
		case "CREATE":
			createCounter.With(prometheus.Labels{"operation": operation}).Inc()
		case "UPDATE":
			updateCounter.With(prometheus.Labels{"operation": operation}).Inc()
		case "DELETE":
			removeCounter.With(prometheus.Labels{"operation": operation}).Inc()
		}

	}
}
