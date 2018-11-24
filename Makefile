.PHONY: all say_hello generate clean

all: build

say_hello:
	@echo "Hello World"

generate:
	@echo "Creating empty text files..."
	touch file-{1..10}.txt

build:
	@echo "Building latest docker image"
	env GOOS=linux GOARCH=amd64 go build --tags netgo --ldflags '-extldflags "-lm -lstdc++ -static"'
	docker build -t monitor:latest .


clean:
	@echo "Cleaning up..."
	rm *.txt