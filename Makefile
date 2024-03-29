# ================================================================================ #
# HELPERS
# ================================================================================ #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ================================================================================ #
# QUALITY CONTROL
# ================================================================================ #

## audit: tidy dependencies and form, vet and test all code
.PHONY: audit
audit: vendor
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...

.PHONY: vendor
vendor:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor

# ================================================================================ #
# BUILD
# ================================================================================ #

## build: builds the srvls binary
.PHONY: build
build: audit
	@echo 'Building cmd/srvls...'
	go build -o=./bin/srvls ./cmd/srvls
	@echo '[OK] Everything is fine so far'
