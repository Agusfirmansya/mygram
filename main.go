package main

import (
	"assignment4_test/routers"
	"assignment4_test/config"
)

func main() {
	config.ConnectDatabase()
	routers.StartServer().Run()
	
}