.PHONY: build rebuild run stop

build:
	docker build -t teslagov/clarakm-projects-go:latest .
	docker-compose build

rebuild:
	docker build -t teslagov/clarakm-projects-go:latest . --no-cache
	docker-compose build

run:
	build
	docker-compose up

stop:
	docker-compose stop