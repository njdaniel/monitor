.PHONY: all say_hello generate clean build deploy

all: build deploy clean

say_hello:
	@echo "Hello World"

generate:
	@echo "Creating empty text files..."
	touch file-{1..10}.txt

build:
	@echo "Building latest docker image"
	env GOOS=linux GOARCH=amd64 go build --tags netgo --ldflags '-extldflags "-lm -lstdc++ -static"'
	docker build -t monitor:latest -f build/Dockerfile .

deploy:
	@echo "Deploying by docker-compose"
	docker-compose -f deployments/docker-compose.yml up


clean:
	@echo "Cleaning up..."
	go clean