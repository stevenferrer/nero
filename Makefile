DOCKER ?= docker

.PHONY: pg-test
pg-test:
	$(DOCKER) rm -f pg-test || true
	$(DOCKER) run --name pg-test -e POSTGRES_PASSWORD=postgres -d --rm -p 5432:5432 postgres:13
	$(DOCKER) exec -it pg-test bash -c 'while ! pg_isready; do sleep 1; done;'