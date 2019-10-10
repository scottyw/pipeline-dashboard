/*
 * START HERE The Compile below com
 */
package report

import (
	"fmt"

	"github.com/puppetlabs/jenkins_report/config"
	"github.com/puppetlabs/jenkins_report/lib/report/cith"
	"github.com/puppetlabs/jenkins_report/lib/report/csv_writers"
	"github.com/puppetlabs/jenkins_report/lib/report/jenkins_types"
	"github.com/puppetlabs/jenkins_report/lib/report/runners"
	"github.com/puppetlabs/jenkins_report/lib/report/utils"
	"strings"
)

func CompileCith(CithURL string) []cith.CithFailure {
	return cith.GetPrevalentFailureCauses(CithURL)
}

func Compile(configData config.Config) []jenkins_types.Pipeline {
	var allPipelines []jenkins_types.Pipeline
	// Pass in kickoff job and follow downstreams

	for _, koJob := range configData.KickoffJobs {
		allPipelines = append(allPipelines, runners.FromKickoff(koJob.URL, koJob.Product, koJob.Version))
	}

	for _, ordJob := range configData.OrderedJobs {
		allPipelines = append(allPipelines, runners.FromOrdered(ordJob.URL, ordJob.Product, ordJob.Version))
	}

	return allPipelines
}

func ApplyCith(pipelines []jenkins_types.Pipeline, cithFailures []cith.CithFailure) []jenkins_types.Pipeline {
	var retVal []jenkins_types.Pipeline

	transients := []string{
		"Infrastructure (Suspected)",
		"Infrastructure (Confirmed)",
	}

	errors := []string{
		"Test (Suspected)",
		"Other",
		"Product (Suspected)",
	}

	fmt.Printf("# of Failures: %d", len(cithFailures))

	for _, pipeline := range pipelines {
		fmt.Println("=============== A Pipeline =============")

		for ji, jobs := range pipeline.TrainData {
			fmt.Printf("# of Jobs: %d\n", len(jobs))
			for jj, job := range jobs {
				for ci, failure := range cithFailures {
					if (job.BuildNumber == failure.BuildNumber || job.BuildNumber-1 == failure.BuildNumber) && strings.Contains(job.URL, failure.Master) && strings.Contains(job.URL, failure.ProjectName) {
						if utils.StringSliceContains(transients, failure.CategoryName) {
							pipeline.Transients++
							pipeline.TrainData[ji][jj].Transients++
							cith.Remove(cithFailures, ci)
						} else if utils.StringSliceContains(errors, failure.CategoryName) {
							pipeline.Errors++
							pipeline.TrainData[ji][jj].Errors++
							cith.Remove(cithFailures, ci)
						} else {
							panic(fmt.Sprintf("%s is not a transient or normal error", failure.CategoryName))
						}
					}
				}
			}
		}
		retVal = append(retVal, pipeline)
	}

	fmt.Println("The following failures did not have any matches: ")
	for _, failure := range cithFailures {
		fmt.Printf("%+v\n", failure)
	}

	csv := jenkins_types.OpenTrainCSV()

	for _, pipeline := range pipelines {
		for _, jobs := range pipeline.TrainData {
			for _, job := range jobs {
				jenkins_types.WriteTrainCSV(csv, job, pipeline.PipelineJob, pipeline.Version)
			}
		}
	}

	return retVal

}

func WriteToCSV(pipelines []jenkins_types.Pipeline) {
	csv_writers.WritePipelines(pipelines)
}
