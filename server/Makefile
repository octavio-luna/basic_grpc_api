postgres:
	docker run --name postgresdb -e POSTGRES_PASSWORD=docker -p 5432:5432 -d postgres
	go run ./db/migration/migration.go
server:
	go build -o bin/server main.go
	echo 'Server built'
	./bin/server
seed_data:
	go run ./sql/seed.go
	echo 'Seed data added'