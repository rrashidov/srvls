# ==========================================================
# Helpers
# ==========================================================
.PHONY: help 
help:
	@echo 'Usage: '
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==========================================================
# Development
# ==========================================================

## build: build the cmd/api application
.PHONY: build
build:
	@echo 'Building cmd/api...'
	go build -o=./bin/api ./cmd/api
	GOOS=linux GOARCH=amd64 go build -o=./bin/linux_amd64/api ./cmd/api
