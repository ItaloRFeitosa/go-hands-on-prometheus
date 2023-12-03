# Get the host user and group IDs
HOST_USER_ID := $(shell id -u)
HOST_GROUP_ID := $(shell id -g)
COMPOSE_USER :=  $(HOST_USER_ID):$(HOST_GROUP_ID)
COMPOSE_FILE := ./deployments/local/docker-compose.yml
.PHONY: run build up

run:
	go run ./cmd/linkapi/main.go

build:
	go build -o ./out/main ./cmd/linkapi/main.go

up:
	COMPOSE_USER=$(COMPOSE_USER) docker-compose --project-directory . -f $(COMPOSE_FILE) up

start_k6:
	COMPOSE_USER=$(COMPOSE_USER) docker-compose --project-directory . -f $(COMPOSE_FILE) start k6

up_load_test:
	COMPOSE_USER=$(COMPOSE_USER) docker-compose --project-directory . -f $(COMPOSE_FILE) --profile loadtest up

stop:
	COMPOSE_USER=$(COMPOSE_USER) docker-compose --project-directory . -f $(COMPOSE_FILE) stop

up_build: 
	COMPOSE_USER=$(COMPOSE_USER) docker-compose --project-directory . -f $(COMPOSE_FILE) up --build

destroy:
	docker-compose --project-directory . -f $(COMPOSE_FILE) down
	sudo rm -rf ./.docker

air_init:
	docker run -it --rm -w /usr/air --user $(COMPOSE_USER) -v $(PWD):/usr/air cosmtrek/air init