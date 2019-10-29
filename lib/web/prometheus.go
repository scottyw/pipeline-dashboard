package web

import (
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func (handlers *Handlers) GenerateMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Our middleware logic goes here...
		for _, job := range handlers.Page.Jobs {
			tM, _ := strconv.Atoi(job.JobDataStrings.TotalMinutes)
			tH, _ := strconv.Atoi(job.JobDataStrings.TotalMinutes)
			wCTM, _ := strconv.Atoi(job.JobDataStrings.WallClockTimeMinutes)
			wCTH, _ := strconv.Atoi(job.JobDataStrings.WallClockTimeMinutes)

			totalSeconds.
				With(prometheus.Labels{
					"pipeline":     job.Pipeline,
					"pipeline_job": job.PipelineJob,
					"version":      job.Version,
					"build_number": strconv.Itoa(job.BuildNumber),
				}).
				Set(float64((60 * tM) + (3600 * tH)))
			wallClockSeconds.
				With(prometheus.Labels{
					"pipeline":     job.Pipeline,
					"pipeline_job": job.PipelineJob,
					"version":      job.Version,
					"build_number": strconv.Itoa(job.BuildNumber),
				}).
				Set(float64((60 * wCTM) + (3600 * wCTH)))
		}

		for _, train := range handlers.Page.Trains {
			totalSeconds.
				With(prometheus.Labels{
					"pipeline":     train.Pipeline,
					"pipeline_job": train.Name,
					"version":      train.Version,
					"build_number": "0",
				}).
				Set(float64(train.DurationSortMinutes * 60))
			wallClockSeconds.
				With(prometheus.Labels{
					"pipeline":     train.Pipeline,
					"pipeline_job": train.Name,
					"version":      train.Version,
					"build_number": "0",
				}).
				Set(float64(train.DurationSortMinutes * 60))
		}

		next.ServeHTTP(w, r)
	})
}

var (
	wallClockSeconds = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "cidashboard_wall_clock_seconds",
		Help: "The time it has taken for a job or pipeline from start to finish.",
	}, []string{"pipeline", "pipeline_job", "version", "build_number"},
	)
	totalSeconds = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "cidashboard_total_seconds",
		Help: "The consecutive time taken for a job or pipeline",
	}, []string{"pipeline", "pipeline_job", "version", "build_number"},
	)
)
