postgersinit:
	docker run --name postgres15 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

postgres:
	docker exec -it postgres15 psql

createDB:
	docker exec -it postgres15 createDB --username=root --owner=root go-chat

dropDB:
	docker exec -it postgres15 dropdb go-chat

.PHONY: postgersinit postgres createDB dropDB
