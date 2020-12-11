PODMAN ?= podman

.PHONY: pg-test
pg-test:
	$(PODMAN) rm -f pg-test || true
	$(PODMAN) run --name pg-test -e POSTGRES_PASSWORD=postgres -d --rm -p 5432:5432 postgres:13
	$(PODMAN) exec -it pg-test bash -c 'while ! pg_isready; do sleep 1; done;'