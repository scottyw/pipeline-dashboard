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

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/puppetlabs/pipeline-dashboard/lib/report/jenkins_types"
)

type Page struct {
	Title    string
	Jobs     []jenkins_types.Pipeline
	Trains   []jenkins_types.TrainStrings
	Products []jenkins_types.Product
}

func (h *Handlers) GeneratePageData() *Page {
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

		errors, _ := strconv.Atoi(line[12])
		transients, _ := strconv.Atoi(line[13])

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
			Errors:     errors,
			Transients: transients,
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

		errors, _ := strconv.Atoi(line[8])
		transients, _ := strconv.Atoi(line[9])

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
			Errors:              errors,
			Transients:          transients,
		})
	}

	sort.Slice(trains, func(i, j int) bool {
		return trains[i].DurationSortMinutes > trains[j].DurationSortMinutes
	})

	for i, product := range h.Products {
		product.SetVals(jobs)
		h.Products[i] = product
	}

	h.Page = &Page{Title: title, Jobs: jobs, Trains: trains, Products: h.Products}

	return h.Page
}

func (handlers *Handlers) IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, handlers.Page)

}

type Handlers struct {
	Products []jenkins_types.Product
	Page     *Page
}

func (handlers *Handlers) ProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(handlers.Page)
}

func Serve() {

	products := jenkins_types.GetProducts()

	fs := http.FileServer(http.Dir("./public/"))

	handlers := &Handlers{Products: products}
	handlers.GeneratePageData()

	http.Handle("/static/css/", fs)
	http.Handle("/static/js/", fs)
	http.Handle("/css/", http.FileServer(http.Dir("./public/")))

	http.Handle("/", http.FileServer(http.Dir("./public/")))
	http.HandleFunc("/api/1/products", handlers.ProductsHandler)

	http.Handle("/metrics", handlers.GenerateMetrics(promhttp.Handler()))

	// http.HandleFunc("/", handler)
	fmt.Println("Listening on port :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
