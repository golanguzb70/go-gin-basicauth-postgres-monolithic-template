run:
	go run cmd/main.go
	
swag_init:
	swag init -g api/router.go  -o api/docs

migrate_up:
	migrate -path migrations/ -database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):5432/$(POSTGRES_DATABASE)?sslmode=disable up

migrate_down:
	migrate -path migrations/ -database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):5432/$(POSTGRES_DATABASE)?sslmode=disable down

migrate_force:
	migrate -path migrations/ -database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):5432/$(POSTGRES_DATABASE)?sslmode=disable force 4

create_migrate:
	./scripts/create_migration.sh

compose_down:  
	docker compose down

compose_up: compose_down
	docker compose up -d --build

crud:
	./scripts/crud.sh