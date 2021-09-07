package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	createCounter     *prometheus.CounterVec
	updateCounter     *prometheus.CounterVec
	removeCounter     *prometheus.CounterVec
	createFailCounter *prometheus.CounterVec
	updateFailCounter *prometheus.CounterVec
	removeFailCounter *prometheus.CounterVec
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

	createCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "createFailCounter"},
		[]string{"operation"})

	updateCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "updateFailCounter"},
		[]string{"operation"})

	removeCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "removeFailCounter"},
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
		case "CREATE_FAIL":
			createFailCounter.With(prometheus.Labels{"operation": operation}).Inc()
		case "UPDATE_FAIL":
			updateFailCounter.With(prometheus.Labels{"operation": operation}).Inc()
		case "DELETE_FAIL":
			removeFailCounter.With(prometheus.Labels{"operation": operation}).Inc()
		}

	}
}
