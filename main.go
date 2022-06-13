package main

import (
	routes "CRUD-Operation/routes"
)

func main() {
	r := routes.StartGin()
	r.Run(":8000")
}
