Set env variable to connect to PG DB

Powershell

$env:PG_CONNSTRING="postgres://username:password@IP:5432/database1"

Linux

export PG_CONNSTRING="postgres://username:password@IP:5432/database1"

Run app

go run main.go

List products
http://localhost:8000/api/product/list-products
