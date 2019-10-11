package jenkins_types

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/puppetlabs/pipeline-dashboard/lib/report/utils"
)

type Train struct {
	JobName         string
	BuildNumber     int
	URL             string
	Name            string
	DurationMinutes float32
	StartTime       time.Time
	EndTime         time.Time
	Timestamp       int64
	Errors          int
	Transients      int
}

type TrainStrings struct {
	URL                 string
	Pipeline            string
	Version             string
	Name                string
	DurationHours       string
	DurationMinutes     string
	DurationSortMinutes int
	StartTime           string
	EndTime             string
	Timestamp           string
	Errors              int
	Transients          int
}

func (t *Train) EndTimeSeconds() int64 {
	return t.Timestamp + (int64(t.DurationMinutes) * 60 * 1000)
}

func (t *Train) StringArray() []string {
	return []string{
		t.URL,
		t.Name,
		fmt.Sprintf("%g", t.DurationMinutes),
		fmt.Sprintf("%s", t.StartTime),
		fmt.Sprintf("%s", t.EndTime),
		fmt.Sprintf("%d", t.Timestamp),
		fmt.Sprintf("%d", t.Errors),
		fmt.Sprintf("%d", t.Transients),
	}
}

func (t *Train) GetEndTime() time.Time {
	retVal := t.StartTime.Add(time.Minute * time.Duration(t.DurationMinutes))
	fmt.Println(t.DurationMinutes)
	fmt.Printf("%s %s\n", t.StartTime, retVal)

	return retVal
}

func OpenTrainCSV() *csv.Writer {
	file, err := os.OpenFile("trains.csv", os.O_WRONLY|os.O_APPEND, 0644)
	utils.CheckError("Cannot create file", err)

	file.WriteString("\n")

	writer := csv.NewWriter(file)
	return writer
}

func WriteTrainCSV(writer *csv.Writer, train Train, jobName string, jobVersion string) {
	writer.Write(append([]string{jobName, jobVersion}, train.StringArray()...))
	writer.Flush()
}
