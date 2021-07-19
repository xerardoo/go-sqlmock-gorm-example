package main

import (
	"github.com/joho/godotenv"
	"github.com/xerardoo/sql-testing-example/models"
	"log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("cannot load env file:", err.Error())
	}

	err = models.InitDB()
	if err != nil {
		log.Fatal("cannot connect to db:", err.Error())
		return
	}
}
