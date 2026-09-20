include .env
export

export PROJECT_ROOT=${shell pwd}

environment-up:
	@docker compose up -d magazine-app-postgres-db

environment-down:
	@docker compose down magazine-app-postgres-db

environment-cleanup:
	@read -p "Do you really wanna clean all volume data(psql container)? Warning:this command cleans all database data! [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down magazine-app-port-forwarder magazine-app-postgres-db && \
		rm -rf out/pgdata && \
		echo "Environment database files was deleted"; \
		else \
		  echo "Close delete environment data"; \
		fi

migrate-create:
	@if [ -z "$(seq)"]; then \
  	echo "Field seq is required. Example: make migrate-create seq={migration_name}" \
  		exit 1; \
  		fi; \
  	docker compose run --rm magazine-app-migrations \
  		create \
  		-ext sql \
  		-dir backend/migrations \
  		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
  echo "Field action is required. Example: make migrate-action action={action}" \
  	exit 1; \
  	fi; \
  docker compose run --rm magazine-app-migrations \
  	  -path /migrations \
  	  -database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@magazine-app-postgres-db:5432/${POSTGRES_DATABASE}?sslmode=disable \
  	  "$(action)"

test-database-data-up:
	@echo "Loading test data to database"
	docker compose exec -T magazine-app-postgres-db \
	psql -U ${POSTGRES_USER} -d ${POSTGRES_DATABASE} < ${PROJECT_ROOT}/backend/data/data.sql


environment-port-forward:
	@docker compose up -d magazine-app-port-forwarder

environment-port-close:
	@docker compose down magazine-app-port-forwarder

magazine-app-backend:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/backend/out/logs && \
	export POSTGRES_HOST=localhost && \
	cd backend && go mod tidy && \
	go run cmd/main.go

magazine-app-frontend:
	@cd frontend && npm run build && serve -s build



