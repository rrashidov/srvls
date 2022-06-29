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

# ==========================================================
# Development environment
# ==========================================================

mysql_root_password=mysqlroot
mysql_port=23306
mysql_db_name=srvlstestdb
mysql_user=srvlstestuser
mysql_password=srvlstestpassword

## setup-test-mysql: run mysql container to use for testing
.PHONY: setup-test-mysql
setup-test-mysql:
	@echo 'Setting up test MySQL instance'
	docker stop srvls-test-mysql || true && docker rm srvls-test-mysql || true
	docker run --detach --name=srvls-test-mysql --env="MYSQL_ROOT_PASSWORD=${mysql_root_password}" --env="MYSQL_DATABASE=${mysql_db_name}" --env="MYSQL_USER=${mysql_user}" --env="MYSQL_PASSWORD=${mysql_password}" --publish ${mysql_port}:3306 mysql:oracle
