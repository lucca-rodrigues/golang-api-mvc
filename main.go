package main

import "github.com/luccarodrigues/golang-api-mvc/server"

func main() {
	s := server.NewServer()

	s.Run()
}