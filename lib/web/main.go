package web

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"encoding/csv"
	"encoding/json"
	"html/template"
	"net/http"
	"sort"
	"strconv"

	"github.com/puppetlabs/jenkins_report/config"
	"github.com/puppetlabs/jenkins_report/lib/report/jenkins_types"
)

type Page struct {
	Title    string
	Jobs     []jenkins_types.Pipeline
	Trains   []jenkins_types.TrainStrings
	Products []config.Product
}

func pageData() *Page {
	title := "CI Dashboard"

	csvFile, _ := os.Open("result.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var jobs []jenkins_types.Pipeline

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		buildNumber, _ := strconv.Atoi(line[5])

		transients, _ := strconv.Atoi(line[12])
		errors, _ := strconv.Atoi(line[13])

		jobs = append(jobs, jenkins_types.Pipeline{
			URL:         line[0],
			Server:      line[1],
			Pipeline:    line[2],
			PipelineJob: line[3],
			Version:     line[4],
			BuildNumber: buildNumber,
			JobDataStrings: &jenkins_types.JobDataStrings{
				StartTime:            line[6],
				EndTime:              line[7],
				WallClockTimeHours:   line[8],
				WallClockTimeMinutes: line[9],
				TotalHours:           line[10],
				TotalMinutes:         line[11],
			},
			Transients: transients,
			Errors:     errors,
		})
	}

	trainCSVFile, _ := os.Open("trains.csv")
	trainReader := csv.NewReader(bufio.NewReader(trainCSVFile))

	var trains []jenkins_types.TrainStrings

	for {
		line, error := trainReader.Read()

		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		minutes, _ := strconv.ParseFloat(line[4], 64)
		hours := minutes / 60
		minutesLeft := int(minutes) % 60

		transients, _ := strconv.Atoi(line[8])
		errors, _ := strconv.Atoi(line[9])

		trains = append(trains, jenkins_types.TrainStrings{
			Pipeline:            line[0],
			Version:             line[1],
			URL:                 line[2],
			Name:                line[3],
			DurationSortMinutes: int(minutes),
			DurationMinutes:     fmt.Sprintf("%d", int(minutesLeft)),
			DurationHours:       fmt.Sprintf("%d", int(hours)),
			StartTime:           line[5],
			EndTime:             line[6],
			Timestamp:           line[7],
			Transients:          transients,
			Errors:              errors,
		})
	}

	sort.Slice(trains, func(i, j int) bool {
		return trains[i].DurationSortMinutes > trains[j].DurationSortMinutes
	})

	products := config.GetProducts()

	for i, product := range products {
		product.SetVals(jobs)
		products[i] = product
	}

	p := &Page{Title: title, Jobs: jobs, Trains: trains, Products: products}

	return p
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p := pageData()

	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, p)

}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	p := pageData()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(p)

}

func Serve() {
	fs := http.FileServer(http.Dir("./public/"))
	http.Handle("/static/css/", fs)
	http.Handle("/static/js/", fs)

	http.Handle("/css/", http.FileServer(http.Dir("./public/")))

	http.Handle("/", http.FileServer(http.Dir("./public/")))

	http.HandleFunc("/api/1/products", productsHandler)

	// http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
