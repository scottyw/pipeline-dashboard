package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"

	"github.com/puppetlabs/jenkins_report/config"
	"github.com/puppetlabs/jenkins_report/lib/report"
	"github.com/puppetlabs/jenkins_report/lib/report/cith"
	"github.com/puppetlabs/jenkins_report/lib/report/jenkins_types"
	"github.com/puppetlabs/jenkins_report/lib/report/utils"
)

func CithFailures(config config.Config) []cith.CithFailure {
	var cithFailures []cith.CithFailure

	if len(config.CithURL) == 0 {
		return cithFailures
	}

	g := utils.Getable{
		URL: "CITHFAILURES",
	}

	client := g.GetRedisClient()
	defer client.Close()
	fmt.Println("Checking to see if CITHFAILURES are cached")

	cached, body := g.Cached(client, "CITHFAILURES")
	json.Unmarshal(body, &cithFailures)

	if cached {
		fmt.Println("CITHFAILURES are cached")
		return cithFailures
	}
	fmt.Println("CITHFAILURES are not cached")

	cithFailures = report.CompileCith(config.CithURL)

	fmt.Printf("Found %d failures for today from Cith.\n", len(cithFailures))
	if len(cithFailures) == 0 {
		panic("Found no Cith Failures")
	}

	body, err := json.Marshal(cithFailures)

	if err != nil {
		panic(err)
	}

	fmt.Println("Caching Cith Failures")
	g.Cache(client, "CITHFAILURES", body)

	return cithFailures
}

func JenkinsData(config config.Config) []jenkins_types.Pipeline {
	var allJenkinsData []jenkins_types.Pipeline

	g := utils.Getable{
		URL: "ALLJENKINSDATA",
	}

	client := g.GetRedisClient()
	defer client.Close()

	cached, body := g.Cached(client, "ALLJENKINSDATA")
	json.Unmarshal(body, &allJenkinsData)

	if cached {
		return allJenkinsData
	}

	allJenkinsData = report.Compile(config)

	body, err := json.Marshal(allJenkinsData)

	if err != nil {
		panic(err)
	}

	g.Cache(client, "ALLJENKINSDATA", body)

	return allJenkinsData
}

func main() {

	file_a, _ := os.Create("result.csv")
	file_b, _ := os.Create("trains.csv")

	writer := csv.NewWriter(file_b)

	writer.Write([]string{
		"URL",
		"ProjectName",
		"ProjectVersion",
		"Name",
		"DurationMinutes",
		"Time",
		"TimeStamp",
	})

	file_a.Close()
	file_b.Close()

	config := config.GetConfig()

	cithFailures := CithFailures(config)
	jenkinsData := JenkinsData(config)

	var compiledData []jenkins_types.Pipeline
	if len(config.CithURL) > 0 {
		compiledData = report.ApplyCith(jenkinsData, cithFailures)
	} else {
		compiledData = jenkinsData
	}

	report.WriteToCSV(compiledData)
}
