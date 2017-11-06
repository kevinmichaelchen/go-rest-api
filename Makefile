.PHONY: pb build rebuild run stop

pb:
	for f in pb/**/*.proto; do \
		protoc --go_out=plugins=grpc:. $$f; \
		echo compiled: $$f; \
	done

build:
	docker build -t teslagov/clarakm-projects-go:latest .
	docker-compose build

rebuild:
	docker build -t teslagov/clarakm-projects-go:latest . --no-cache
	docker-compose build

run:
	docker-compose up

stop:
	docker-compose stop