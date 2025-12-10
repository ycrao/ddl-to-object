.PHONY: build clean test fmt vet install

# Version information
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
GIT_COMMIT ?= $(shell git rev-parse HEAD 2>/dev/null || echo "unknown")
BUILD_TIME ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

# Build flags
LDFLAGS = -X 'ddl-to-object/cmd.Version=$(VERSION)' \
          -X 'ddl-to-object/cmd.GitCommit=$(GIT_COMMIT)' \
          -X 'ddl-to-object/cmd.BuildTime=$(BUILD_TIME)' \
		  -s -w

# Build the binary
build:
	mkdir -p bin
	GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o bin/ddl-to-object .
	cp -r template/ bin/template
	cp -r config.example.json bin/config.json

# Clean build artifacts
clean:
	rm -rf bin/
	rm -rf release/

# Run tests
test:
	go test ./...

# Format code
fmt:
	go fmt ./...

# Vet code
vet:
	go vet ./...

# Install dependencies
deps:
	go mod tidy
	go mod download

# Install binary to GOPATH/bin
install:
	go install .

# Build for multiple platforms
build-all: clean
	mkdir -p release
	GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o release/linux/ddl-to-object .
	GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o release/mac/ddl-to-object .
	GOOS=darwin GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o release/mac-arm64/ddl-to-object .
	GOOS=windows GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o release/win/ddl-to-object.exe .
	cp -r template/ release/linux/template
	cp -r template/ release/mac/template
	cp -r template/ release/mac-arm64/template
	cp -r template/ release/win/template
	cp -r config.example.json release/linux/config.json
	cp -r config.example.json release/mac/config.json
	cp -r config.example.json release/mac-arm64/config.json
	cp -r config.example.json release/win/config.json
