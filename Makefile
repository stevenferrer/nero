DOCKER ?= docker

.PHONY: postgres-test
postgres-test:
	$(DOCKER) rm -f postgres-test || true
	$(DOCKER) run --name postgres-test -e POSTGRES_PASSWORD=postgres -d --rm -p 5432:5432 postgres:12
	$(DOCKER) exec -it postgres-test bash -c 'while ! pg_isready; do sleep 1; done;'