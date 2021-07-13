.DEFAULT_GOAL := help
.PHONY: help up down build rebuild logs go

go-image := golang:1.16.6-buster
docker-dir := docker
docker-compose := docker-compose --env-file=$(docker-dir)/.env -f $(docker-dir)/docker-compose.yml
green-color := \033[32m
reset-color := \033[0m

help:
	@echo "Run $(green-color)make up$(reset-color) to set up project. All orphan containers will be removed."
up:
	@make down
	@make build
	$(docker-compose) up -d --remove-orphans
	docker ps

down:
	$(docker-compose) down

build:
	$(docker-compose) build

rebuild:
	$(docker-compose) build --pull --no-cache

logs:
	$(docker-compose) logs -f

go-build:
	docker run --rm -v "$$PWD"/go-app:/go/src/app -w /go/src/app $(go-image) go build -v

go-run:
	docker run --rm -v "$$PWD"/go-app:/go/src/app -w /go/src/app $(go-image) go run .