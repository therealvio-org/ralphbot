bin:
	mkdir ./bin

clean:
	rm -rf ./bin

.PHONY: mod
mod:
	go mod download

.PHONY: build
build: bin mod
	GOOS=linux CGO_ENABLED=0 go build -v -o bin/ralphbot ./cmd/ralphbot

.PHONY: build-deployment-binary
build-fargate: bin mod
	GOARCH=arm64 GOOS=linux CGO_ENABLED=0 go build -v -o bin/ralphbot ./cmd/ralphbot

.PHONY: local
local: bin mod build
	docker-compose -f docker-compose-local.yaml up --build