###Install node_exporter

Download node_exporter

	$ wget https://github.com/prometheus/node_exporter/releases/download/v0.16.0/node_exporter-0.16.0.linux-amd64.tar.gz
	$ tar -xf node_exporter-0.16.0.linux-amd64.tar.gz
	$ sudo mv node_exporter-0.16.0.linux-amd64/node_exporter /usr/local/bin
	$ rm -r node_exporter-0.16.0.linux-amd64*

	$ sudo useradd -rs /bin/false node_exporter
	
/etc/systemd/system/node_exporter.service

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
	
Start daemon

	$ sudo systemctl daemon-reload
	$ sudo systemctl enable node_exporter
    $ sudo systemctl start node_exporter

###Install Prometheus

	$ wget https://github.com/prometheus/prometheus/releases/download/v2.5.0/prometheus-2.5.0.linux-amd64.tar.gz
	$ tar -xf prometheus-2.5.0.linux-amd64.tar.gz
    $ sudo mv prometheus-2.5.0.linux-amd64/prometheus prometheus-2.5.0.linux-amd64/promtool /usr/local/bin
    $ sudo mkdir /etc/prometheus /var/lib/prometheus
	$ sudo mv prometheus-2.5.0.linux-amd64/consoles prometheus-2.5.0.linux-amd64/console_libraries /etc/prometheus
	$ rm -r prometheus-2.5.0-linux-amd64*
	
/etc/prometheus/prometheus.yml

	global:
	  scrape_interval: 10s
	
	scrape_configs:
	  - job_name: 'prometheus_metrics'
		scrape_interval: 5s
		static_configs:
		  - targets: ['localhost:9090']
	  - job_name: 'node_exporter_metrics'
		scrape_interval: 5s
		static_configs:
		  - targets: ['localhost:9100']

Change ownership
  
  	sudo useradd -rs /bin/false prometheus
	sudo chown -R prometheus: /etc/prometheus /var/lib/prometheus
	
/etc/systemd/system/prometheus.service

	[Unit]
	Description=Prometheus
	After=network.target
	
	[Service]
	User=prometheus
	Group=prometheus
	Type=simple
	ExecStart=/usr/local/bin/prometheus \
		--config.file /etc/prometheus/prometheus.yml \
		--storage.tsdb.path /var/lib/prometheus/ \
		--web.console.templates=/etc/prometheus/consoles \
		--web.console.libraries=/etc/prometheus/console_libraries
	
	[Install]
	WantedBy=multi-user.target
	
Start service
	
	sudo systemctl daemon-reload
	sudo systemctl enable prometheus
    sudo systemctl start prometheus