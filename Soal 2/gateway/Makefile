all: build

export GO111MODULE	?= on

dep:
	GO111MODULE=on go mod download

tidy:
	go mod tidy
	gofmt -s -w ./**/*.go

build: dep
	GO111MODULE=on CGO_ENABLED=0 go build -ldflags="-s -w" \
	-o build/bibit-gateway .

docker-build:
	docker build -t bibit-gateway:1.0.0 .

docker-up:
	docker-compose up

run: build
	./build/bibit-gateway

test:
	go test -v ./... -cover -vet -all -coverprofile=coverage.out

cover:
	go tool cover -html=coverage.out

generate:
	go generate ./ent