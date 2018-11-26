#!/usr/bin/env bash

wget https://github.com/prometheus/node_exporter/releases/download/v0.16.0/node_exporter-0.16.0.linux-amd64.tar.gz

tar -xf node_exporter-0.16.0.linux-amd64.tar.gz

sudo mv node_exporter-0.16.0.linux-amd64/node_exporter /usr/local/bin

rm -r node_exporter-0.16.0.linux-amd64*

sudo useradd -rs /bin/false node_exporter

#Create service
if [[ ! -e /etc/systemd/system/node_exporter.service ]]; then
        sudo touch /etc/systemd/system/node_exporter.service
fi
cat > /etc/systemd/system/node_exporter.service << EOF1
[Unit]
Description=Node Exporter
After=network.target

[Service]
User=node_exporter
Group=node_exporter
Type=simple
ExecStart=/usr/local/bin/node_exporter

[Install]
WantedBy=multi-user.target

EOF1

sudo systemctl daemon-reload
sudo systemctl enable node_exporter
sudo systemctl start node_exporter
