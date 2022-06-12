Simple golang crud that creates, lists, deletes and updates a list of products

Set env variable to connect to PG DB

Powershell

$env:PG_CONNSTRING="postgres://username:password@IP:5432/database1"

Linux

export PG_CONNSTRING="postgres://username:password@IP:5432/database1"

Run app

go run main.go

List products
http://localhost:8000/api/product/list-products

Create migration schema files
migrate create -ext sql -dir db/migration -seq <name_prefix>