package main

import (
	"crud-go/db"
	"crud-go/routes"
)

func main() {
	e := routes.Init()

	db.Init()

	e.Logger.Fatal(e.Start(":9000"))
}
