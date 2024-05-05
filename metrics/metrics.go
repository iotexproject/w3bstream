package metrics

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	dispatchedTaskNumMtc = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "dispatched_task_num_metrics",
			Help: "dispatched task num metrics.",
		}, []string{"projectID", "projectVersion"})
	retryTaskNumMtc = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "retry_task_num_metrics",
			Help: "retry task num metrics.",
		}, []string{"projectID", "projectVersion"})
	timeoutTaskNumMtc = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "timeout_task_num_metrics",
			Help: "timeout task num metrics.",
		}, []string{"projectID", "projectVersion"})
	taskDurationMtc = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "task_duration_metrics",
		Help: "task duration metrics.",
	}, []string{"projectID", "projectVersion"})
	taskRuntimeMtc = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "task_runtime_metrics",
		Help:    "task runtime metrics.",
		Buckets: prometheus.LinearBuckets(0, 60, 10),
	}, []string{"projectID", "projectVersion"})
	failedTaskNumMtc = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "failed_task_num_metrics",
			Help: "failed task num metrics.",
		}, []string{"projectID", "projectVersion"})
	succeedTaskNumMtc = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "succeed_task_num_metrics",
			Help: "succeed task num metrics.",
		}, []string{"projectID", "projectVersion"})
	taskFinalStateNumMtc = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "task_final_state_num_metrics",
		Help: "task final state num metrics.",
	}, []string{"projectID", "projectVersion", "state"})
)

func init() {
	prometheus.MustRegister(dispatchedTaskNumMtc)
	prometheus.MustRegister(retryTaskNumMtc)
	prometheus.MustRegister(timeoutTaskNumMtc)
	prometheus.MustRegister(taskDurationMtc)
	prometheus.MustRegister(failedTaskNumMtc)
	prometheus.MustRegister(succeedTaskNumMtc)
	prometheus.MustRegister(taskFinalStateNumMtc)
	prometheus.MustRegister(taskRuntimeMtc)
}

func DispatchedTaskNumMtc(projectID uint64, projectVersion string) {
	dispatchedTaskNumMtc.WithLabelValues(strconv.FormatUint(projectID, 10), projectVersion).Inc()
}

func RetryTaskNumMtc(projectID uint64, projectVersion string) {
	retryTaskNumMtc.WithLabelValues(strconv.FormatUint(projectID, 10), projectVersion).Inc()
}

func TimeoutTaskNumMtc(projectID uint64, projectVersion string) {
	timeoutTaskNumMtc.WithLabelValues(strconv.FormatUint(projectID, 10), projectVersion).Inc()
}

func TaskDurationMtc(projectID uint64, projectVersion string, duration float64) {
	taskDurationMtc.WithLabelValues(strconv.FormatUint(projectID, 10), projectVersion).Set(duration)

	taskRuntimeMtc.WithLabelValues(strconv.FormatUint(projectID, 10), projectVersion).Observe(duration)
}

func FailedTaskNumMtc(projectID uint64, projectVersion string) {
	failedTaskNumMtc.WithLabelValues(strconv.FormatUint(projectID, 10), projectVersion).Inc()
}

func SucceedTaskNumMtc(projectID uint64, projectVersion string) {
	succeedTaskNumMtc.WithLabelValues(strconv.FormatUint(projectID, 10), projectVersion).Inc()
}

func TaskFinalStateNumMtc(projectID uint64, projectVersion, state string) {
	taskFinalStateNumMtc.WithLabelValues(strconv.FormatUint(projectID, 10), projectVersion, state).Inc()
}
