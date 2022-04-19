.PHONY: build
build:
	docker build . --tag theikalman/product:latest

.PHONY: run
run:
	docker-compose up --detach

.PHONY: stop
stop:
	docker-compose down