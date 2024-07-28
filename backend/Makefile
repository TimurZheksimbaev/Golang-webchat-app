run:
	@go run cmd/main.go

source:
	@source app.env

migrateup:
	@migrate -path database/migrations -database $(DB_URL) -verbose up

migratedown: 
	@migrate -path database/migrations -database $DB_URL -verbose down