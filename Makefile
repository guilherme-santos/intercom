.PHONY: all goget test clean install run

BINARY=invite-customers
VERSION=1.0.0

all: test ${BINARY}

goget:
	go get ./...

test: goget
	go test ./...

${BINARY}:
	go build -o ${BINARY} main.go

clean:
	go clean
	rm -rf ${BINARY}

install: all
	go install .

run: all
	./${BINARY}