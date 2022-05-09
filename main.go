package main

import (
	"digimer-api/src/database"
	"digimer-api/src/routes"
)

func init() {
	new(database.DBConf).InitDB().Migrate()
}

func main() {
	app := routes.New()
	app.Logger.Fatal(app.Start(":8000"))
}
