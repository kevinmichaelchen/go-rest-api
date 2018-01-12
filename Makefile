.PHONY: pb
.PHONY: build rebuild
.PHONY: start stop seed

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

start:
	docker-compose up

stop:
	docker-compose stop

seed:
	./seed-data.sh