package main

import (
	"NOMOR1/config"
	"NOMOR1/routes"
)

func main() {
	config.Init()
	e := routes.New()
	e.Start(":8001")
}