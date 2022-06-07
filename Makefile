createdb:
	docker exec -it postgres createdb --username=root --owner=root products
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/products?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/products?sslmode=disable" -verbose down

.PHONY: createdb migrateup migratedown