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