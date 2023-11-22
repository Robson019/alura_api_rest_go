package main

import (
	"fmt"

	"github.com/Robson019/api-rest-go/database"
	"github.com/Robson019/api-rest-go/routes"
)

func main() {
	database.Connection()
	fmt.Println("Starting Rest server with Go")
	fmt.Println("http://localhost:8000")
	routes.HandleRequest()
}
