include .env
export

export PROJECT_ROOT=${shell pwd}

environment-up:
	@docker compose up -d magazine-app-postgres-db

environment-down:
	@docker compose down magazine-app-postgres-db

