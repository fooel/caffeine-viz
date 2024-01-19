BINARY_PATH=bin/caffeine-viz

all: build

build:
	go build -o "${BINARY_PATH}" main.go

run:
	go build -o "${BINARY_PATH}" main.go 
	./"${BINARY_PATH}"

clean:
	go clean
	rm "${BINARY_PATH}"
