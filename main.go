package main

import (
	"fmt"

	"github.com/marcosmourabro/go-rest-api/database"
	"github.com/marcosmourabro/go-rest-api/routes"
)

func main() {

	database.ConnectionWithDatabase()
	fmt.Println("Starting Rest server with Go")
	routes.HandleRequest()
}
