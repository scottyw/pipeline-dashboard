package jenkins_types

import (
	"fmt"
	"github.com/puppetlabs/pipeline-dashboard/lib/report/utils"
	"strings"
	"time"
)

type Builds struct {
	List []Build
}

func ProcessTopLevelBuilds(jd BuildsAndJobs) Builds {
	/*
	 * First grab all top level jobs.
	 */
	var retVal Builds

	for i, build := range jd.Builds {
		utils.Log(fmt.Sprintf("Processing Build #%d: %s\n", i, build.FullDisplayName), build.URL)
		jobForBuild := FindJob(build, jd)

		if build.URL == "" {
			fmt.Println("There was no URL associated with this build.")
			continue
		}

		jobForBuild.Fetch()
		build.Fetch()

		// var buildDownstreamJobs []Job
		var buildDownstreamBuilds []Build

		if len(jobForBuild.DownstreamProjects) > 0 {
			buildDownstreamBuilds, _ = JobsFromDownstreamProjects(build, jobForBuild.DownstreamProjects)
		}

		build.JobName = jobForBuild.Name
		retVal.List = append(retVal.List, build)
		retVal.List = append(retVal.List, buildDownstreamBuilds...)
		// for n, job := range buildDownstreamJobs {
		// 	fmt.Printf("!%d", n)
		// 	lastBuild := job.LastBuild
		//
		// 	if strings.Contains(lastBuild.URL, "http") {
		// 		lastBuild.Fetch()
		//
		// 		build.JobName = jobForBuild.Name
		//
		// 		if BuildTriggerMatchesParent(lastBuild, build) {
		// 			retVal.List = append(retVal.List, lastBuild)
		// 			utils.Log(
		// 				fmt.Sprintf("Job: %s || Build has number: %d", jobForBuild.Name, build.Number),
		// 				build.URL,
		// 			)
		// 		}
		// 	} else {
		// 		continue
		// 	}
		// }
	}
	return retVal
}

/* AllBuildsTriggerMatchesParent iterates through builds struct of a job and check
if any matches the parent */
func AllBuildsTriggerMatchesParent(child Job, parent Build) (bool, Build) {
	utils.LogHeading(fmt.Sprintf("Checking all builds for %s to see if %s is the parent", child.Name, parent.FullDisplayName), "")
	fmt.Printf("--- %s\n", child.Name)

	for _, build := range child.Builds {
		utils.LogTree(fmt.Sprintf("Trying %s", build.FullDisplayName), "", 1)
		if BuildTriggerMatchesParent(build, parent) {
			build.Fetch()
			utils.LogTree(fmt.Sprintf("Returning Build with Number %d", build.Number), "", 3)
			return true, build
		}
	}

	var build Build

	return false, build
}

func BuildTriggerMatchesParent(child Build, parent Build) bool {
	triggered_url_stub, triggered_by := child.TriggeredBy()
	if triggered_by == 0 {
		child.Fetch()
		triggered_url_stub, triggered_by = child.TriggeredBy()
	}

	if parent.Number == 0 || parent.URL == "" {
		parent.Fetch()
	}

	if parent.URL == "" {
		panic("No Parent URL")
	}
	utils.LogTree(fmt.Sprintf("Child was triggered by %d, parent is %d (%s should contain %s)", triggered_by, parent.Number, parent.URL, triggered_url_stub), "", 2)

	return (triggered_by == parent.Number) && strings.Contains(parent.URL, triggered_url_stub)
}

func LogTree(trainData map[int][]Train) {

	utils.LogHeading("Train Tree: ", "")
	var jobs []string
	for _, topTrain := range trainData {
		for _, train := range topTrain {
			if train.JobName != "" {
				jobs = append(jobs, train.JobName)
			}
		}
	}

	for _, job := range jobs {
		utils.LogTree(job, "", 1)
		for _, topTrain := range trainData {
			for _, train := range topTrain {
				fmt.Printf("%s %s", train.JobName, job)
				if train.JobName == job {
					utils.LogTree(fmt.Sprintf("%s %d %s %s", train.Name, train.BuildNumber, train.StartTime, train.EndTime), "", 2)
				}
			}
		}
	}
}

func (b *Builds) GetJobData(pipeline_name string, pipeline_version string) (JobData, map[int][]Train) {
	trainData := make(map[int][]Train)

	for _, build := range b.List {
		build.Fetch()
		i := 0
		fmt.Printf("%s\n", build.FullDisplayName)
		fmt.Printf("%s\n", build.URL)
		fmt.Printf("Processing Build %s\n\n", build.Class)

		if build.Class == "hudson.matrix.MatrixBuild" {
			// Here is where we get trains from Matrix Builds

			if len(build.Runs) > 0 {

				for _, cell_build := range BuildsFromMatrixRuns(build, build.Runs) {

					var train Train
					utils.Log(fmt.Sprintf("Matrix Cell: %s", cell_build.FullDisplayName), cell_build.URL)

					train.JobName = cell_build.JobName
					train.BuildNumber = cell_build.Number
					train.URL = cell_build.URL
					train.Name = cell_build.FullDisplayName
					train.DurationMinutes = float32(cell_build.Duration) / (60 * 1000)
					train.StartTime = time.Unix(cell_build.Timestamp/1000, 0)
					train.EndTime = train.GetEndTime()
					train.Timestamp = cell_build.Timestamp

					if len(trainData[i]) > 0 {
						trainData[i] = append(trainData[i], train)
					} else {
						trainData[i] = []Train{train}
					}
				}
			}

			i++
		} else {
			var train Train

			train.JobName = build.JobName
			train.BuildNumber = build.Number
			train.Name = build.FullDisplayName
			train.URL = build.URL
			train.DurationMinutes = float32(build.Duration) / (60 * 1000)
			train.StartTime = time.Unix(build.Timestamp/1000, 0)
			train.EndTime = train.GetEndTime()
			train.Timestamp = build.Timestamp

			if len(trainData[i]) > 0 {
				trainData[i] = append(trainData[i], train)
			} else {
				trainData[i] = []Train{train}
			}
			i++

		}

	}

	utils.LogHeading(fmt.Sprintf("Processing %d Trains\n\n", len(trainData)), "")
	utils.LogHeading(fmt.Sprintf("Processing %d Trains\n\n", len(trainData[0])), "")

	var jobData JobData

	LogTree(trainData)

	for _, train := range trainData {
		var totalMinutes float32
		totalMinutes = 0

		var startTime int64
		var endTime int64

		startTime = 9999999999999
		endTime = 0

		for _, t := range train {
			totalMinutes = totalMinutes + t.DurationMinutes

			timeOfEvent := time.Unix(t.Timestamp/1000, 0)

			if time.Now().Sub(timeOfEvent).Hours()/24 <= 365 {

				timestamp := time.Unix(t.Timestamp/1000, 0)

				if startTime > t.Timestamp {
					startTime = t.Timestamp
					fmt.Printf("TRAIN UPDATE START %s started at %d-%02d-%02dT%02d:%02d:%02d-00:00\n", t.Name, timestamp.Year(), timestamp.Month(), timestamp.Day(), timestamp.Hour(), timestamp.Minute(), timestamp.Second())
				}

				if endTime < t.EndTimeSeconds() {
					endTime = t.EndTimeSeconds()
					fmt.Printf("TRAIN UPDATE END %s ended at %d-%02d-%02dT%02d:%02d:%02d-00:00\n", t.Name, timestamp.Year(), timestamp.Month(), timestamp.Day(), timestamp.Hour(), timestamp.Minute(), timestamp.Second())
				}
			} else {
				// fmt.Printf("Skipping Job because %s is more than 7 days ago\n", timeOfEvent)
			}

		}

		jobData.AssignJobValues(startTime, endTime, totalMinutes)
	}

	utils.LogHeading("Pipeline Data", "")
	utils.Log(fmt.Sprintf("Start Time: %s", jobData.StartTime), "")
	utils.Log(fmt.Sprintf("End Time: %s", jobData.EndTime), "")
	utils.Log(fmt.Sprintf("Wall Clock Time Hours: %d", jobData.WallClockTimeHours), "")
	utils.Log(fmt.Sprintf("Wall Clock Time Minutes: %d", jobData.WallClockTimeMinutes), "")

	return jobData, trainData

}
