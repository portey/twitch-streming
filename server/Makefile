.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -o artifacts/svc .

.PHONY: deps
deps:
	go mod tidy
	go mod download
	go mod vendor