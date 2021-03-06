DIST_PATH ?= ./dist
MAIN ?= main.go
NAME ?= disco

build:
	go build -o ${DIST_PATH}/${NAME}${EXT} ${MAIN}
.PHONY: build

build-all:
	make clean

	EXT=-darwin-386 GOARCH=386 GOOS=darwin make build
	EXT=-darwin-amd64 GOARCH=amd64 GOOS=darwin make build

	EXT=-linux-arm GOARCH=arm GOOS=linux make build
	EXT=-linux-arm64 GOARCH=arm64 GOOS=linux make build
	EXT=-linux-amd64 GOARCH=amd64 GOOS=linux make build
	EXT=-linux-386 GOARCH=386 GOOS=linux make build

	EXT=-win-386.exe GOARCH=386 GOOS=windows make build
	EXT=-win-amd64.exe GOARCH=amd64 GOOS=windows make build
.PHONY: build-all

clean:
	rm -Rf ${DIST_PATH}
.PHONY: clean

install:
	dep ensure -v
.PHONY: install

run:
	go run ${MAIN} ${ARGS}
.PHONY: run