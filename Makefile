DOCKER ?= docker

.PHONY: postgres
postgres:
	$(DOCKER) rm -f postgres || true
	$(DOCKER) run --name postgres -e POSTGRES_PASSWORD=postgres -d --rm -p 5432:5432 postgres:13
	$(DOCKER) exec -it postgres bash -c 'while ! pg_isready; do sleep 1; done;'