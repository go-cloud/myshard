all: build

build: build-proxy build-config build-agent

build-proxy:
	go build -o bin/my-proxy ./cmd/proxy

build-config:
	go build -o bin/my-config ./cmd/config

build-agent:
	go build -o bin/my-agent ./cmd/agent

clean:
	@rm -rf bin

test:
	go test ./... -race
