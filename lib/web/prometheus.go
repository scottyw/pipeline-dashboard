package web

import (
	"fmt"

	"net/http"
	"regexp"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func (handlers *Handlers) GenerateMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Our middleware logic goes here...
		for _, job := range handlers.Page.Jobs {
			tH, _ := strconv.Atoi(job.JobDataStrings.TotalHours)
			tM, _ := strconv.Atoi(job.JobDataStrings.TotalMinutes)
			wCTH, _ := strconv.Atoi(job.JobDataStrings.WallClockTimeHours)
			wCTM, _ := strconv.Atoi(job.JobDataStrings.WallClockTimeMinutes)

			totalSeconds.
				With(prometheus.Labels{
					"pipeline":     job.Pipeline,
					"pipeline_job": job.PipelineJob,
					"version":      job.Version,
					"build_number": strconv.Itoa(job.BuildNumber),
					"scope":        "pipeline",
				}).
				Set(float64((60 * tM) + (3600 * tH)))
			wallClockSeconds.
				With(prometheus.Labels{
					"pipeline":     job.Pipeline,
					"pipeline_job": job.PipelineJob,
					"version":      job.Version,
					"build_number": strconv.Itoa(job.BuildNumber),
					"scope":        "pipeline",
				}).
				Set(float64((60 * wCTM) + (3600 * wCTH)))
		}

		for _, train := range handlers.Page.Trains {
			var matchBuildNumber = regexp.MustCompile(` #[0-9]+$`)

			var trainNameWithoutBuildNumber = string(matchBuildNumber.ReplaceAll([]byte(train.Name), []byte("")))
			fmt.Println(trainNameWithoutBuildNumber)

			totalSeconds.
				With(prometheus.Labels{
					"pipeline":     train.Pipeline,
					"pipeline_job": trainNameWithoutBuildNumber,
					"version":      train.Version,
					"build_number": "0",
					"scope":        "job",
				}).
				Set(float64(train.DurationSortMinutes * 60))
			wallClockSeconds.
				With(prometheus.Labels{
					"pipeline":     train.Pipeline,
					"pipeline_job": trainNameWithoutBuildNumber,
					"version":      train.Version,
					"build_number": "0",
					"scope":        "job",
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
	}, []string{"pipeline", "pipeline_job", "version", "build_number", "scope"},
	)
	totalSeconds = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "cidashboard_total_seconds",
		Help: "The consecutive time taken for a job or pipeline",
	}, []string{"pipeline", "pipeline_job", "version", "build_number", "scope"},
	)
)
