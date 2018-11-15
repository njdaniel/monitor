#Monitor

Tutorial
    https://blog.alexellis.io/prometheus-monitoring/
    https://github.com/alexellis/hash-browns

Build

    env GOOS=linux GOARCH=amd64 go build --tags netgo --ldflags '-extldflags "-lm -lstdc++ -static"'

Start prometheus server

    $ docker-compose up

ALTERNATIVE dockerstack

    $ hostname -I
    $ docker swarm init --advertise-addr 192.168.1.124:2377