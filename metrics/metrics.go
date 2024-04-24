package metrics

import (
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

func TaskStartTimeMtc(projectID, projectVersion, taskID string) {
	taskStartTimeMtc.WithLabelValues(projectID, projectVersion, taskID).SetToCurrentTime()
}

func DispatchedTaskNumMtc(projectID, projectVersion string) {
	dispatchedTaskNumMtc.WithLabelValues(projectID, projectVersion).Inc()
}

func RetryTaskNumMtc(projectID, projectVersion, taskID string) {
	retryTaskNumMtc.WithLabelValues(projectID, projectVersion, taskID).Inc()
}

func TimeoutTaskNumMtc(projectID, projectVersion, taskID string) {
	timeoutTaskNumMtc.WithLabelValues(projectID, projectVersion, taskID).Inc()
}

func TaskEndTimeMtc(projectID, projectVersion, taskID string) {
	taskEndTimeMtc.WithLabelValues(projectID, projectVersion, taskID).SetToCurrentTime()
}

func FailedTaskNumMtc(projectID, projectVersion string) {
	failedTaskNumMtc.WithLabelValues(projectID, projectVersion).Inc()
}

func SucceedTaskNumMtc(projectID, projectVersion string) {
	succeedTaskNumMtc.WithLabelValues(projectID, projectVersion).Inc()
}

func TaskFinalStateMtc(projectID, projectVersion, taskID, state string) {
	taskFinalStateMtc.WithLabelValues(projectID, projectVersion, taskID, state).SetToCurrentTime()
}
