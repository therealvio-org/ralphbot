bin:
	mkdir ./bin

clean:
	rm -rf ./bin

.PHONY: mod
mod:
	go mod download

deps: mod
	go mod tidy
	git diff --exit-code

.PHONY: build
build: bin mod
	GOOS=linux CGO_ENABLED=0 go build -v -o bin/ralphbot .
#CGO_ENABLED=0 to force a static build. Next stage of build uses an image that has libraries in different locations
#https://stackoverflow.com/a/36308464

.PHONY: build-deployment-binary
build-deployment-binary: bin mod
	GOARCH=arm64 GOOS=linux CGO_ENABLED=0 go build -v -o bin/ralphbot .

.PHONY: local
local: bin mod build
	docker-compose -f docker-compose-local.yaml up --build