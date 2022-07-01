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

mysql_container_name=srvls-test-db

## setup-test-mysql: run mysql container to use for testing
.PHONY: setup-test-mysql
setup-test-mysql:
	@echo 'Setting up test MySQL instance'
	@echo 'Stopping any old instances'
	@docker stop ${mysql_container_name} || true && docker rm ${mysql_container_name} || true
	@echo 'Old test instance stopped'
	@echo 'Run new MySQL container'
	@docker run --detach --name=${mysql_container_name} --health-cmd='mysqladmin ping --silent' --env="MYSQL_ROOT_PASSWORD=${mysql_root_password}" --env="MYSQL_DATABASE=${mysql_db_name}" --env="MYSQL_USER=${mysql_user}" --env="MYSQL_PASSWORD=${mysql_password}" --publish ${mysql_port}:3306 mysql:oracle
	@echo 'New MySQL container created. Wait for it to start'
	@while [ "`docker inspect -f {{.State.Health.Status}} ${mysql_container_name}`" != "healthy" ]; do sleep 2; done
	@echo 'New MySQL container started. Create tables'
	@migrate -path=./migrations -database="mysql://${mysql_user}:${mysql_password}@tcp(127.0.0.1:${mysql_port})/${mysql_db_name}" up
	@echo 'Tables created. DB is ready for testing'
