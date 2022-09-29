package main

import (
	"github.com/luccarodrigues/golang-api-mvc/database"
	"github.com/luccarodrigues/golang-api-mvc/server"
)

func main() {
	database.StartDB()
	s := server.NewServer()

	s.Run()
}