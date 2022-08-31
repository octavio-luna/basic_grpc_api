
# Basic gRPC API 

This is a simple implementation of a RPC server and
client. It's build in golang with a default postgres database.



## How to build and run it

### If you have make installed 
To build and run server, type in a new terminal 
in root folder:
    
    make server

To run migration:

    make migration

To load seed data:

    make seed

### NOTE: If you don't have make installed 
To build and run server, type in a new terminal 
in root folder:

    docker-compose build
	docker-compose up
To run migration:

    docker exec server go run server/db/migration/migration.go

To load seed data:

    docker exec server go run server/sql/seed.go


## CLIENT

The /client/ folder contains a basic script for testing the 
server endpoints.

### If you have make installed 
To run the client type in a new terminal:

    make client

### NOTE: If you don't have make installed 
To run the client type in a new terminal:

    docker exec server go run client/main.go

## IMPORTANT
To run the client you should have made the migration and added 
the seed data before, otherway, you'll come to many database-related
problems