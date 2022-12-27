postgres:
	docker run --name postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:14-alpine


createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres  web_events

dropdb:
	docker exec -it postgres dropdb web_events

sqlc:
	sqlc generate

.PHONY: sqlc dropdb