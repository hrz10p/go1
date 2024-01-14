run:
	@go run ./cmd/

migrate:
	@sqlite3 db.sqlite < ./pkg/migrations/tables.sql
	@echo "Migrated tables"

reloader:
	@./rungo.sh ./cmd/

