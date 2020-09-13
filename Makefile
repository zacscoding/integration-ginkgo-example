MODULE = $(shell go list -m)

.PHONY: build build-docker compose compose-down
build: # build a server
	go build -a -o integration-server $(MODULE)/cmd/server

build-docker: # build docker image
	docker build -f cmd/server/Dockerfile -t integration/integration-server .

compose: # run with docker-compose
	docker-compose up --force-recreate

compose-down: # down docker-compose
	docker-compose down -v