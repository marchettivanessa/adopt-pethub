ROOT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
OS := $(shell uname -s)
export PATH := bin:$(PATH)

start:
	docker-compose up -d

status:
	docker-compose ps

stop:
	docker-compose down

restart:
	docker-compose down && docker-compose up -d && make status

install-mockgen:
	go install github.com/golang/mock/mockgen@73266f9

generate-mocks:
	mockgen -source=handler/handler.go -destination=handler/mocks/mock_interface.go -package=mocks

run-adopt-pethub:
	set -o allexport && source .env && go run adopt-pethub/main.go
