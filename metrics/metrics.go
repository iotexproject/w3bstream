package metrics

import (
	"strconv"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var timeMap sync.Map

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
	taskEndTimeMtc = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "task_end_time_metrics",
		Help: "task end time metrics.",
	}, []string{"projectID", "projectVersion"})
	taskRuntimeMtc = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "task_runtime_metrics",
		Help:    "task runtime metrics.",
		Buckets: nil,
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
	prometheus.MustRegister(taskEndTimeMtc)
	prometheus.MustRegister(failedTaskNumMtc)
	prometheus.MustRegister(succeedTaskNumMtc)
	prometheus.MustRegister(taskFinalStateNumMtc)
	prometheus.MustRegister(taskRuntimeMtc)
}

func TaskStartTimeMtc(taskID uint64) {
	timeMap.Store(taskID, float64(time.Now().UnixNano())/1e9)
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

func TaskEndTimeMtc(projectID, taskID uint64, projectVersion string) {
	start, _ := timeMap.Load(taskID)
	duration := float64(time.Now().UnixNano())/1e9 - start.(float64)
	timeMap.Delete(taskID)
	taskEndTimeMtc.WithLabelValues(strconv.FormatUint(projectID, 10), projectVersion).Set(duration)

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
