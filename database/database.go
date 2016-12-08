package database

import (
	"github.com/Dunkelheit/feedbackapp/model"
	"github.com/jinzhu/gorm"
)

// DB is the database instance
var DB *gorm.DB

// OpenDB opens the database
func OpenDB() {
	var err error
	DB, err = gorm.Open("postgres", "sslmode=disable dbname=feedbackapp host=localhost user=arturo.martinez")
	if err != nil {
		panic("Failed to connect database")
	}
	// defer DB.Close()

	DB.LogMode(true)

	DB.DropTable(&model.Review{}, &model.Card{}, &model.User{}, &model.Role{})
	DB.AutoMigrate(&model.Card{}, &model.User{}, &model.Review{}, &model.Role{})

	DB.Model(&model.Review{}).AddForeignKey("reviewer_id", "users(id)", "CASCADE", "CASCADE")
	DB.Model(&model.Review{}).AddForeignKey("reviewee_id", "users(id)", "CASCADE", "CASCADE")

	DB.Create(&model.Card{Title: "Great moves", Category: model.CardCategoryPositive})
	DB.Create(&model.Card{Title: "Computer hacking skills", Category: model.CardCategoryPositive})
	DB.Create(&model.Card{Title: "Lazy eye", Category: model.CardCategoryNegative})
	DB.Create(&model.Card{Title: "Stinky feet", Category: model.CardCategoryNegative})
}
