.PHONY: build
build:
	docker build . --tag theikalman/user:latest

.PHONY: run
run:
	docker-compose up --detach

.PHONY: stop
stop:
	docker-compose down