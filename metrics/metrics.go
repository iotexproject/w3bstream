package metrics

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	taskStartTimeMtc = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "task_start_time_metrics",
		Help: "task start time metrics.",
	}, []string{"projectID", "projectVersion", "taskID"})
	dispatchedTaskNumMtc = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "dispatched_task_num_metrics",
			Help: "dispatched task num metrics.",
		}, []string{"projectID", "projectVersion"})
	retryTaskNumMtc = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "retry_task_num_metrics",
			Help: "retry task num metrics.",
		}, []string{"projectID", "projectVersion", "taskID"})
	timeoutTaskNumMtc = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "timeout_task_num_metrics",
			Help: "timeout task num metrics.",
		}, []string{"projectID", "projectVersion", "taskID"})
	taskEndTimeMtc = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "task_end_time_metrics",
		Help: "task end time metrics.",
	}, []string{"projectID", "projectVersion", "taskID"})
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
	taskFinalStateMtc = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "task_final_state_metrics",
		Help: "task final state metrics.",
	}, []string{"projectID", "projectVersion", "taskID", "state"})
)

func init() {
	prometheus.MustRegister(taskStartTimeMtc)
	prometheus.MustRegister(dispatchedTaskNumMtc)
	prometheus.MustRegister(retryTaskNumMtc)
	prometheus.MustRegister(timeoutTaskNumMtc)
	prometheus.MustRegister(taskEndTimeMtc)
	prometheus.MustRegister(failedTaskNumMtc)
	prometheus.MustRegister(succeedTaskNumMtc)
	prometheus.MustRegister(taskFinalStateMtc)
}

func TaskStartTimeMtc(projectID, taskID uint64, projectVersion string) {
	taskStartTimeMtc.WithLabelValues(strconv.FormatUint(projectID, 10), projectVersion, strconv.FormatUint(taskID, 10)).SetToCurrentTime()
}

func DispatchedTaskNumMtc(projectID uint64, projectVersion string) {
	dispatchedTaskNumMtc.WithLabelValues(strconv.FormatUint(projectID, 10), projectVersion).Inc()
}

func RetryTaskNumMtc(projectID, taskID uint64, projectVersion string) {
	retryTaskNumMtc.WithLabelValues(strconv.FormatUint(projectID, 10), projectVersion, strconv.FormatUint(taskID, 10)).Inc()
}

func TimeoutTaskNumMtc(projectID, taskID uint64, projectVersion string) {
	timeoutTaskNumMtc.WithLabelValues(strconv.FormatUint(projectID, 10), projectVersion, strconv.FormatUint(taskID, 10)).Inc()
}

func TaskEndTimeMtc(projectID, taskID uint64, projectVersion string) {
	taskEndTimeMtc.WithLabelValues(strconv.FormatUint(projectID, 10), projectVersion, strconv.FormatUint(taskID, 10)).SetToCurrentTime()
}

func FailedTaskNumMtc(projectID uint64, projectVersion string) {
	failedTaskNumMtc.WithLabelValues(strconv.FormatUint(projectID, 10), projectVersion).Inc()
}

func SucceedTaskNumMtc(projectID uint64, projectVersion string) {
	succeedTaskNumMtc.WithLabelValues(strconv.FormatUint(projectID, 10), projectVersion).Inc()
}

func TaskFinalStateMtc(projectID, taskID uint64, projectVersion, state string) {
	taskFinalStateMtc.WithLabelValues(strconv.FormatUint(projectID, 10), projectVersion, strconv.FormatUint(taskID, 10), state).SetToCurrentTime()
}
