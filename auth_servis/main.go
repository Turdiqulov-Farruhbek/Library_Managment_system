package main

import (
	"fmt"
	"log"

	"github.com/Project_Restaurant/Auth-Service/api"
	"github.com/Project_Restaurant/Auth-Service/postgres"
	_ "github.com/Project_Restaurant/Auth-Service/docs"
)

func main() {
	db, err := postgres.ConnectDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	router := api.NewRouter(db)
	
	fmt.Println("Server is running on port 2121")
	if err := router.Run(":2121"); err != nil {
		log.Fatal(err)
	}
	
}
