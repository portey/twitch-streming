.PHONY: gen
gen:
	go get github.com/99designs/gqlgen
	go run -mod=mod github.com/99designs/gqlgen generate

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -o artifacts/svc .

.PHONY: deps
deps:
	go mod tidy
	go mod download
	go mod vendor