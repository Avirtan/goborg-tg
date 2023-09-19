include .env
export

compose-up: ### Run docker-compose
	docker-compose up --build -d 
.PHONY: compose-up

migrate-create:  ### create new migration
	migrate create -ext sql -dir migrations 'migrate_name'
.PHONY: migrate-create

migrate-up: ### migration up
	migrate -path migrations -database '$(PG_URL)?sslmode=disable' up
.PHONY: migrate-up

migrate-down: ### migration up
	migrate -path migrations -database '$(PG_URL)?sslmode=disable' down
.PHONY: migrate-up

run:
	go run cmd/app/main.go
.PHONY: run