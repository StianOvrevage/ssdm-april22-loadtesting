echo "Starting influxdb and grafana"
sudo influxd & sudo grafana-server -homepath /usr/share/grafana/ -config /etc/grafana/grafana.ini
