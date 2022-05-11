package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sync"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

func main() {

	var metrics Metrics

	// Read source file
	file, err := os.Open("data/results.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// InfluxDB connection
	client := influxdb2.NewClientWithOptions("http://localhost:8086", "",
		influxdb2.DefaultOptions().SetBatchSize(50000))

	defer client.Close()

	writeAPI := client.WriteAPI("", "k6")

	//q := client.NewQuery("CREATE DATABASE k6", "", "")
	//if response, err := c.Query(q); err == nil && response.Error() == nil {
	//	fmt.Printf("Create DB results: %+v\n", response.Results)
	//}

	// Create channel for points feeding
	pointsCh := make(chan *write.Point, 20000)

	threads := 5

	var writerWg sync.WaitGroup
	// Launch write routines
	for t := 0; t < threads; t++ {
		writerWg.Add(1)
		go func() {
			for p := range pointsCh {
				writeAPI.WritePoint(p)
			}
			writerWg.Done()
		}()
	}

	// Processing
	pointRegexp, _ := regexp.Compile("{\"type\":\"Point\"")
	metricRegexp, _ := regexp.Compile("{\"type\":\"Metric\"")

	var line string

	var p *write.Point
	for scanner.Scan() {
		line = scanner.Text()

		if pointRegexp.MatchString(line) {
			var point K6Point
			if err := json.Unmarshal([]byte(line), &point); err != nil {
				panic(err)
			}

			if point.Metric != "http_req_duration" && point.Metric != "vus" && point.Metric != "http_req_failed" {
				metrics.SkippedWrongMetric++
				continue
			}

			//fmt.Printf("Queueing %v : %v : %v\n", point.Data.Time.Format("2006-01-02T15:04:05.999999"), point.Metric, point.Data.Value)

			p = influxdb2.NewPoint(
				point.Metric,
				map[string]string{
					"url": point.Data.Tags.URL,
				},
				map[string]interface{}{
					"value": point.Data.Value,
				},
				point.Data.Time,
			)

			metrics.PointsBatched++
			pointsCh <- p
		}

		if metricRegexp.MatchString(line) {
			metrics.SkippedMetricNotPoint++
			continue

			var metric K6Metric
			if err := json.Unmarshal([]byte(line), &metric); err != nil {
				panic(err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	close(pointsCh)
	writeAPI.Flush()

	// Wait for writes complete
	writerWg.Wait()
	fmt.Printf("Metrics: %+v\n", metrics)
}

// {
// 	"type":"Metric",
// 		"data":{
// 		"name":"http_req_tls_handshaking",
// 		"type":"trend",
// 		"contains":"time",
// 		"tainted":null,
// 		"thresholds":[],
// 		"submetrics":null,
// 		"sub":{
// 			"name":"",
// 			"parent":"",
// 			"suffix":"",
// 			"tags":null
// 		}
// 	},
// 	"metric":"http_req_tls_handshaking"
// }

type K6Metric struct {
	Type string `json:"type"`
	Data struct {
		Name       string        `json:"name"`
		Type       string        `json:"type"`
		Contains   string        `json:"contains"`
		Tainted    bool          `json:"tainted"`
		Thresholds []interface{} `json:"thresholds"`
		Submetrics interface{}   `json:"submetrics"`
		Sub        struct {
			Name   string      `json:"name"`
			Parent string      `json:"parent"`
			Suffix string      `json:"suffix"`
			Tags   interface{} `json:"tags"`
		} `json:"sub"`
		Metric string `json:"metric"`
	} `json:"data"`
}

// {
// 	"type":"Point",
// 	"data":{
// 		"time":"2022-05-11T18:30:09.6175791+02:00",
// 		"value":19.1331,
// 		"tags":{
// 			"expected_response":"true",
// 			"group":"",
// 			"method":"GET",
// 			"name":"https://loadtest.signicat.dev/health",
// 			"proto":"HTTP/2.0",
// 			"scenario":"default",
// 			"status":"200",
// 			"tls_version":"tls1.3",
// 			"url":"https://loadtest.signicat.dev/health"
// 		}
// 	},
// 	"metric":"http_req_connecting"
// }

type K6Point struct {
	Type string `json:"type"`
	Data struct {
		Time  time.Time `json:"time"`
		Value float32   `json:"value"`
		Tags  struct {
			ExpectedResponse string `json:"expected_response"` // Or bool?
			Group            string `json:"group"`
			Method           string `json:"method"`
			Name             string `json:"name"`
			Proto            string `json:"proto"`
			Scenario         string `json:"scenario"`
			Status           string `json:"status"` // Or int?
			TLSVersion       string `json:"tls_version"`
			URL              string `json:"url"`
		} `json:"tags"`
	} `json:"data"`
	Metric string `json:"metric"`
}

type Metrics struct {
	SkippedMetricNotPoint int
	SkippedWrongMetric    int
	PointsBatched         int
}
