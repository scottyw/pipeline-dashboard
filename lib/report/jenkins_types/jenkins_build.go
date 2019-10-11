package jenkins_types

import (
	"encoding/json"
	"fmt"
	"github.com/puppetlabs/pipeline-dashboard/lib/report/utils"

	"strings"
)

type Build struct {
	utils.Getable
	JobName         string
	Actions         []Action  `json:"actions"`
	Class           string    `json:"_class"`
	Number          int       `json:"number"`
	Duration        int       `json:"duration"`
	Name            string    `json:"name"`
	FullDisplayName string    `json:"fullDisplayName"`
	Runs            []Project `json:"runs"`
	Timestamp       int64     `json:"timestamp"`
	URL             string    `json:"url"`
	TopLevelBuild   bool
	TimeInQueue     TimeInQueue
}

func TriggeredBuildsFromActions(actions []Action) []Build {
	var retVal []Build
	for _, a := range actions {
		if a.Class == "hudson.plugins.parameterizedtrigger.BuildInfoExporterAction" {
			retVal = a.TriggeredBuilds
		}
	}

	return retVal
}

func BuildsFromMatrixRuns(parent Build, projects []Project) []Build {
	var builds []Build
	for _, dp := range projects {

		build := Build{URL: dp.URL}
		build.Fetch()

		if strings.Contains(build.FullDisplayName, "g6002a5f") {
			fmt.Printf("Parent: %s #%d\n", parent.FullDisplayName, parent.Number)
			fmt.Printf("Child: %s #%d\n", build.FullDisplayName, build.Number)
		}

		if BuildTriggerMatchesParent(build, parent) {
			builds = append(builds, build)
		}

	}

	return builds
}

func (b *Build) filter() string {
	return "api/json?tree=actions[triggeredBuilds[*],causes[*]],duration,fullDisplayName,name,number,timestamp,url,runs[number,url]&depth=2"
}

func (b *Build) TriggeredBy() (string, int) {
	// var project string
	var buildNumber int
	var url string
	for _, action := range b.Actions {
		if action.Class == "hudson.model.CauseAction" {
			// project = action.Causes[0].UpstreamProject
			buildNumber = action.Causes[0].UpstreamBuild
			url = action.Causes[0].UpstreamURL

			return url, buildNumber
		}
	}

	return url, buildNumber
}

func (b *Build) Fetch() Build {

	urlWithFilter := fmt.Sprint(b.URL + b.filter())

	if b.URL == "" {
		fmt.Println("ERROR build.go line 42: b.URL not set")
	} else {
		body := b.Get(urlWithFilter)
		json.Unmarshal(body, &b)
	}

	return *b
}
