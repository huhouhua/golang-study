GO := go
.DEFAULT_GOAL := tidy

.PHONY: start_docker
start_docker:
	docker compose up

.PHONY: down_docker
down_docker:
	docker compose down

.PHONY: tidy
tidy:
	@$(GO) mod tidy