package main

import (
	"github.com/DanielDDHM/management/config"
	"github.com/DanielDDHM/management/database"
	"github.com/DanielDDHM/management/server"
)

func main() {
	config.Init()
	database.StartDatabase()
	server := server.NewServer()
	server.Run()
}
