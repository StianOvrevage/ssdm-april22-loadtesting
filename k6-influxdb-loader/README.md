# k6-influxdb-loader

Tool to load k6.io JSON metrics into InfluxDB.

Created to solve scaling problems where for example InfluxDB is not able to ingest data fast enough. A typical problem is 413 Request too large.

It also enables larger load tests on a single machine since metrics ingestion is not done at the same time competing for resources.

It uses v2 of the InfluxDB Client and is compatible with both InfluxDB 1.8 and 2.x (but not tested on 2.x).

## Installation

Install golang. Install dependencies:

    go mod tidy

## Usage

Run a test outputing metrics to JSON file:

    k6 run --no-summary --out json=data/results.json script.js

Load metrics into influxdb running on `localhost:8086` into database `k6`:

    go run loader.go

> If you want to change filename or InfluxDB hostname, authentication etc you need to edit `loader.go` yourself before running :)

Currently only `http_req_duration`, `vus` and `http_req_failed` metrics will be inserted.

## Performance

A k6 test with 6 million requests running for 6 minutes resulted in ~19GB of metrics.

k6-influxdb-loader inserted three of the metrics `http_req_duration`, `vus` and `http_req_failed` in 8 minutes.

> Both loader and InfluxDB has very low CPU usage but I strongly suspect InfluxDB is the bottleneck and that the naive
concurrency I've copied from the documentation is not for increasing performance but simply to enable multiple write sources.

## Backlog

 - gzip compression for metrics json file to save disk space
 - Custom JSON decoder
 - Metrics
    - Runtime
    - Time waiting for InfluxDB
 - Configurability:
   - InfluxDB database settings
   - Metrics to include/exclude
   - Tags to include/exclude
