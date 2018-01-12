.PHONY: all
.PHONY: test
.PHONY: build rebuild
.PHONY: start remove stop
.PHONY: seed
.PHONY: list-users create-user

all:
	@$(MAKE) remove
	@$(MAKE) rebuild
	@$(MAKE) start
	sleep 10
	@$(MAKE) seed

test:
	go test -v

build:
	docker build -t teslagov/clarakm-projects-go:latest .
	docker-compose build

rebuild:
	docker build -t teslagov/clarakm-projects-go:latest . --no-cache
	docker-compose build

remove:
	docker rm --force clarakm-projects-go || true

start:
	docker-compose up -d

stop:
	docker-compose stop

seed:
	./seed-data.sh

list-users:
	http localhost:8080/users

create-user:
	curl -X POST localhost:8080/user -d '{"name":"Kevin","age":24}'