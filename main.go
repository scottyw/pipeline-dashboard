package main

import (
	"encoding/csv"
	"os"

	"github.com/puppetlabs/jenkins_report/config"
	"github.com/puppetlabs/jenkins_report/lib/report"
)

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

	report.Compile(config)
}
