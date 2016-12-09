package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/Dunkelheit/feedbackgame-backend/database"
	"github.com/Dunkelheit/feedbackgame-backend/router"
)

func main() {
	router.NewRouter().Run()
}

func init() {
	database.OpenDB()
}
