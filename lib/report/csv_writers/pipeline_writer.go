package csv_writers

import (
	"encoding/csv"
	"github.com/puppetlabs/pipeline-dashboard/lib/report/jenkins_types"
	"github.com/puppetlabs/pipeline-dashboard/lib/report/utils"
	"os"
)

func WritePipelines(allPipelines []jenkins_types.Pipeline) {
	file, err := os.OpenFile("result.csv", os.O_WRONLY|os.O_APPEND, 0644)
	utils.CheckError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{
		"URL",
		"Server",
		"Pipeline",
		"Pipeline Job",
		"Version",
		"Build Number",
		"Start Time",
		"End Time",
		"Wall Clock Time Hours",
		"Wall Clock Time Minutes",
		"Total Hours",
		"Total Minutes",
		"Errors",
		"Transients",
	})

	for _, value := range allPipelines {
		err := writer.Write(value.StringArray())
		utils.CheckError("Cannot write to file", err)
	}
}
