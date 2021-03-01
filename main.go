package main

import (
	databases "test_mekar_api/database"
	"test_mekar_api/routes"
)

func main() {
	databases.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":8888"))
}
