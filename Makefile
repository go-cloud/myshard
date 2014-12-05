all: build

build: build-proxy build-config build-keeper

build-proxy:
	go build -o bin/my-proxy ./cmd/proxy

build-config:
	go build -o bin/my-config ./cmd/config

build-keeper:
	go build -o bin/my-keeper ./cmd/keeper

clean:
	@rm -rf bin

test:
	go test ./... -race
