/*
 * START HERE The Compile below com
 */
package report

import (
	"github.com/puppetlabs/jenkins_report/config"
	"github.com/puppetlabs/jenkins_report/lib/report/csv_writers"
	"github.com/puppetlabs/jenkins_report/lib/report/jenkins_types"
	"github.com/puppetlabs/jenkins_report/lib/report/runners"
	"time"
)

func Compile(configData config.Config) {
	var allPipelines []jenkins_types.Pipeline
	// Pass in kickoff job and follow downstreams

	for _, koJob := range configData.KickoffJobs {
		allPipelines = append(allPipelines, runners.FromKickoff(koJob.URL, koJob.Product, koJob.Version))
		time.Sleep(time.Second)
	}

	for _, ordJob := range configData.OrderedJobs {
		allPipelines = append(allPipelines, runners.FromOrdered(ordJob.URL, ordJob.Product, ordJob.Version))
		time.Sleep(time.Second)
	}

	csv_writers.WritePipelines(allPipelines)
}
