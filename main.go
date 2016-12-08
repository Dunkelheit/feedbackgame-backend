package main

import (
	"github.com/Dunkelheit/feedbackgame-backend/database"
	"github.com/Dunkelheit/feedbackgame-backend/router"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	router.NewRouter().Run()
}

func init() {
	database.OpenDB()
}
