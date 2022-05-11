echo "Starting influxdb and grafana"

export INFLUXDB_DATA_QUERY_LOG_ENABLED=false
export INFLUXDB_HTTP_PPROF_ENABLED=false
export INFLUXDB_HTTP_MAX_BODY_SIZE=250000000
export INFLUXDB_CONTINUOUS_QUERIES_ENABLED=false

influxd & grafana-server -homepath /usr/share/grafana/ -config /etc/grafana/grafana.ini
