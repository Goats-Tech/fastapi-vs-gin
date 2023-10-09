new-migration:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrate-up:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/gin?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/gin?sslmode=disable" -verbose down

sqlc:
	sqlc generate