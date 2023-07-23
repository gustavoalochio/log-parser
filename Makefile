ifndef $(GOPATH)
  export GOPATH := $(shell go env GOPATH)
  export PATH := $(GOPATH)/bin:$(PATH)
endif

build:
	@go build -o log-parser main.go

start: build
	@./log-parser

mock-gen:
	@go generate ./...

test:
	go test ./... -coverprofile=cover.out ./... && go tool cover -html=cover.out && make show-cov

show-cov:
	go tool cover -func ./cover.out | tail -n 1 | xargs -n1 | tail -n 1

install-mock-gen:
	@go install github.com/golang/mock/mockgen@latest