package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	SyncedBlockHeightMtc = prometheus.NewGauge(prometheus.GaugeOpts{ //
		Name: "synced_block_height",
		Help: "Height of the latest synced block.",
	})
	ProverMtc = prometheus.NewGauge(prometheus.GaugeOpts{ //
		Name: "prover_total",
		Help: "Total number of provers.",
	})
	NewTaskMtc = prometheus.NewCounterVec( //
		prometheus.CounterOpts{
			Name: "new_tasks_total",
			Help: "Total number of new tasks.",
		}, []string{"projectID"})
	AssignedTaskMtc = prometheus.NewCounterVec( //
		prometheus.CounterOpts{
			Name: "assigned_tasks_total",
			Help: "Total number of tasks that have been assigned.",
		}, []string{"projectID"})
	FailedAssignedTaskMtc = prometheus.NewCounterVec( //
		prometheus.CounterOpts{
			Name: "failed_assigned_tasks_total",
			Help: "Total number of tasks that have failed to be assigned.",
		}, []string{"projectID"})
	TaskDurationMtc = prometheus.NewGaugeVec(prometheus.GaugeOpts{ //
		Name: "task_duration_seconds",
		Help: "Duration of task execution in seconds.",
	}, []string{"projectID", "projectVersion", "taskID"})
	FailedTaskNumMtc = prometheus.NewCounterVec( //
		prometheus.CounterOpts{
			Name: "failed_tasks_total",
			Help: "Total number of tasks that have failed.",
		}, []string{"projectID"})
	SucceedTaskNumMtc = prometheus.NewCounterVec( //
		prometheus.CounterOpts{
			Name: "successful_tasks_total",
			Help: "Total number of tasks that have completed successfully.",
		}, []string{"projectID"})
)

func Init() {
	prometheus.MustRegister(SyncedBlockHeightMtc)
	prometheus.MustRegister(ProverMtc)
	prometheus.MustRegister(NewTaskMtc)
	prometheus.MustRegister(AssignedTaskMtc)
	prometheus.MustRegister(FailedAssignedTaskMtc)
	prometheus.MustRegister(TaskDurationMtc)
	prometheus.MustRegister(FailedTaskNumMtc)
	prometheus.MustRegister(SucceedTaskNumMtc)
}

// RegisterMetrics adds prometheus metrics endpoint to gin engine
func RegisterMetrics(r *gin.Engine) {
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
