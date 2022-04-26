
sudo apt-get install -y apt-transport-https
sudo apt-get install -y software-properties-common wget

echo "Installing locust"
sudo -H pip3 install locust

echo "Installing golang. AND removing old versions."
wget https://go.dev/dl/go1.18.1.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.18.1.linux-amd64.tar.gz
rm go1.18.1.linux-amd64.tar.gz

export PATH=$PATH:/usr/local/go/bin


echo "Installing k6"

sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys C5AD17C747E3415A3642D57D77C6C491D6AC1D69
echo "deb https://dl.k6.io/deb stable main" | sudo tee /etc/apt/sources.list.d/k6.list
sudo apt-get update && sudo apt-get install -y k6


echo "Installing influxdb 1.8"

wget -qO- https://repos.influxdata.com/influxdb.key | gpg --dearmor | sudo tee /etc/apt/trusted.gpg.d/influxdb.gpg
export DISTRIB_ID=$(lsb_release -si); export DISTRIB_CODENAME=$(lsb_release -sc)
echo "deb [signed-by=/etc/apt/trusted.gpg.d/influxdb.gpg] https://repos.influxdata.com/${DISTRIB_ID,,} ${DISTRIB_CODENAME} stable" | sudo tee /etc/apt/sources.list.d/influxdb.list
sudo apt-get update && sudo apt-get install -y influxdb


echo "Installing grafana"

wget -q -O - https://packages.grafana.com/gpg.key | sudo apt-key add -
echo "deb https://packages.grafana.com/oss/deb stable main" | sudo tee -a /etc/apt/sources.list.d/grafana.list
sudo apt-get update && sudo apt-get install -y grafana

echo "min_refresh_interval = 1s" | sudo tee -a /etc/grafana/grafana.ini

echo "Downloading influxdb data source and dashboard provisioning for grafana"
sudo wget -O /usr/share/grafana/conf/provisioning/datasources/grafana-datasource-influxdb.yaml https://raw.githubusercontent.com/StianOvrevage/ssdm-april22-loadtesting/main/grafana-datasource-influxdb.yaml
sudo wget -O /usr/share/grafana/conf/provisioning/dashboards/grafana-dashboard-provisioning.yaml https://raw.githubusercontent.com/StianOvrevage/ssdm-april22-loadtesting/main/grafana-dashboard-provisioning.yaml

# Doesn't work :(
#echo "Downloading dashboard"
#sudo apt-get install -y jq
#sudo mkdir -p /var/lib/grafana/dashboards
#wget -O - https://grafana.com/api/dashboards/4411/revisions/4/download | jq '.id |= "1"' | sudo tee /var/lib/grafana/dashboards/k6-load-testing-results.json

# Uninstalling and deleting data:
#
# sudo apt-get remove -y grafana
# sudo rm -rf /etc/grafana/
# sudo rm -rf /usr/share/grafana
# sudo rm -rf /var/lib/grafana
# sudo rm -rf /var/log/grafana 
# 
# sudo apt-get remove -y influxdb
# sudo rm -rf /etc/influxdb
# sudo rm -rf /var/lib/influxdb