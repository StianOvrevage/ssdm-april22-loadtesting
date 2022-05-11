# ssdm-april22-loadtesting
Load testing - Stavanger Software Developers Meetup April 2022

## Run test

Benchmark max performance:

    k6 run --no-summary --vus 10 --duration 20s script-advanced.js

Load test and send metrics to InfluxDB:

    k6 run --out influxdb=http://localhost:8086/myk6db --duration 20s script-advanced.js

Load test and send metrics to InfluxDB. Do not summarize to conserve resources while running:

    k6 run --no-summary --out influxdb=http://localhost:8086/myk6db --vus 10 --duration 20s script-advanced.js

## Tweaking performance of k6 on Windows Subsystem for Linux (WSL2)

    sudo sysctl -w net.ipv4.ip_local_port_range="1024 65535"
    sudo sysctl -w net.ipv4.tcp_tw_reuse=1
    sudo sysctl -w net.ipv4.tcp_timestamps=1
    sudo prlimit --nofile=250000 --pid $$; ulimit -n 250000

## Installing tools

### k6

Ref: https://k6.io/docs/getting-started/installation/

    sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys C5AD17C747E3415A3642D57D77C6C491D6AC1D69
    echo "deb https://dl.k6.io/deb stable main" | sudo tee /etc/apt/sources.list.d/k6.list
    sudo apt-get update
    sudo apt-get install -y k6

### influxDB

PS: k6 does not support influxdb2 natively because of backwards compatibility.

Ref: https://docs.influxdata.com/influxdb/v1.8/introduction/install/

    wget -qO- https://repos.influxdata.com/influxdb.key | gpg --dearmor > /etc/apt/trusted.gpg.d/influxdb.gpg
    export DISTRIB_ID=$(lsb_release -si); export DISTRIB_CODENAME=$(lsb_release -sc)
    echo "deb [signed-by=/etc/apt/trusted.gpg.d/influxdb.gpg] https://repos.influxdata.com/${DISTRIB_ID,,} ${DISTRIB_CODENAME} stable" | sudo tee /etc/apt/sources.list.d/influxdb.list

    sudo apt-get update && sudo apt-get install -y influxdb

Starting influxdb:

    sudo influxdb &

### Grafana

    sudo apt-get install -y apt-transport-https
    sudo apt-get install -y software-properties-common wget
    wget -q -O - https://packages.grafana.com/gpg.key | sudo apt-key add -
    echo "deb https://packages.grafana.com/oss/deb stable main" | sudo tee -a /etc/apt/sources.list.d/grafana.list
    sudo apt-get update && sudo apt-get install -y grafana
    echo "min_refresh_interval = 1s" | sudo tee -a /etc/grafana/grafana.ini

Starting grafana:
    
    sudo grafana-server -config /etc/grafana/grafana.ini -homepath /usr/share/grafana/ &

Login http://localhost:3000 with admin/admin and import the K6 load testing dashboard with id 4411.
