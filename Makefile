# Makefile for gdsu

.PHONY: all vet staticcheck tidy test fmt lint ci

# Run everything (safe defaults)
all: fmt vet staticcheck test

# Format all code
fmt:
	go fmt ./...

# Run go vet
vet:
	go vet ./...

# Run staticcheck (must be installed with: go install honnef.co/go/tools/cmd/staticcheck@latest)
staticcheck:
	staticcheck ./...

# Run unit tests
test:
	go test -v ./...

# Keep go.mod and go.sum tidy
tidy:
	go mod tidy

# Lint: run all local checks
lint: fmt vet staticcheck tidy

# CI helper: runs everything except tidy (used in GitHub Actions)
ci: fmt vet staticcheck test

coverage:
	go test -coverprofile=coverage.out ./...

covhtml: coverage
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html

bench:
	go test -bench=. ./...

clean:
	rm -f coverage.out
	rm -f coverage.html
