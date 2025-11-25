# Makefile for gdsu

.PHONY: all vet staticcheck tidy test fmt lint ci coverage covhtml bench doc clean install

install:
	go install honnef.co/go/tools/cmd/staticcheck@2025.1.1
	go install golang.org/x/tools/cmd/godoc@latest

# Run everything (safe defaults)
all: fmt vet staticcheck test

# Format all code
fmt:
	go fmt ./...

# Run go vet
vet:
	go vet ./...

# Run static analysis
staticcheck: install
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
	@printf "View file://${PWD}/coverage.html#file0 \n"

bench:
	go test -bench=. ./...

doc: install
	@printf "View http://localhost:6060/pkg/github.com/arunksaha/gdsu/, CTRL-C when done\n"
	godoc -http=:6060

clean:
	rm -f coverage.out
	rm -f coverage.html
