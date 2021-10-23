package main

import (
	"restfulapi/ridzqi/config"
	"restfulapi/ridzqi/routes"
)

func main() {
	config.InitDB()
	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))
}
