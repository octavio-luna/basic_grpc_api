.PHONY: server migration seed client

server:
	docker-compose build
	docker-compose up

migration:
	docker exec server go run server/db/migration/migration.go

seed:
	docker exec server go run server/sql/seed.go

client:
	docker exec server go run client/main.go