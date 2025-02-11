postgres:
	docker run --name event-db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=securepassword -d postgres:12-alpine

createDb:
	docker exec -it event-db createdb --username=root --owner=root event-ticketing

dropDb:
	docker exec -it event-db dropdb event-ticketing

migrateUp:
	migrate -path db/migration -database="postgresql://root:securepassword@localhost:5432/event-ticketing?sslmode=disable" -verbose up

migrateDown:
	migrate -path db/migration -database="postgresql://root:securepassword@localhost:5432/event-ticketing?sslmode=disable" -verbose down


.PHONY: postgres createDb dropDb migrateUp migrateDown