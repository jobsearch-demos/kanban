# Language: makefile
# Path: Makefile

run:
	docker-compose -f ./deployment/docker/docker-compose.yml up
build:
	docker-compose -f ./deployment/docker/docker-compose.yml build
run-prod:
	docker-compose -f ./deployment/docker/docker-compose.yml -f ./deployment/docker/docker-compose.prod.yml up

build-prod:
	docker-compose -f ./deployment/docker/docker-compose.yml -f ./deployment/docker/docker-compose.prod.yml build

# TODO: // Fill in the rest of the commands
