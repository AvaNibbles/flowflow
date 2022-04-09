export GO111MODULE=on

build:
	go build -v ./cmd/flowflow
test:
	go test -v -race -short ./...
lint:
	golangci-lint run
dev:
	air

# Formatting
fmt: fmt-swag fmt-mod-tidy fmt-imports fmt-fmt
fmt-imports:
	goimports -w .
fmt-fmt:
	gofmt -w -s .
fmt-mod-tidy:
	go mod tidy
fmt-swag:
	swag fmt

# Codegen
gen: gen-swagger

gen-swagger:
	rm -rf ./pkg/docs ./pkg/client ./pkg/models
	swag init --output ./pkg/docs
	swagger generate client -f .\pkg\docs\swagger.json -t ./pkg -A flowflowclient