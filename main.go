package main

import "digimer-api/src/routes"

func main() {
	app := routes.New()
	app.Logger.Fatal(app.Start(":8000"))
}
