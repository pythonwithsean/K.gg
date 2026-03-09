.SILENT:


run:
	go run cmd/main.go

build:
	go build -o bin/out cmd/main.go

clean:
	rm -rf ./bin
