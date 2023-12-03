GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest

.PHONY: proto
proto:
	@protoc --proto_path=. --micro_out=./ --go_out=:./  proto/*.proto
.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build -o bin/api_bin api/main.go
	@go build -o bin/backend_bin backend/main.go
	@go build -o bin/gomicro_gin_http_api gomicro_gin_http_api/main.go
	@go build -o bin/gomicro_web_http_api /gomicro_web_http/main.go

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: docker
docker:
	# @docker build -t api:latest .
