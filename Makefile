run_service:
	go run ./cmd/service/main.go --config config/config.yml

create_migration:
	migrate create -ext sql -dir migrations -seq $(name)
