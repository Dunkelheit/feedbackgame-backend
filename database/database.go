package database

import (
	"github.com/Dunkelheit/feedbackgame-backend/model"
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

	DB.Create(&model.Epic{Title: "Visionair"})
	DB.Create(&model.Epic{Title: "Social"})
	DB.Create(&model.Epic{Title: "Hands On"})
	DB.Create(&model.Epic{Title: "Grounded"})
	DB.Create(&model.Epic{Title: "Brainpower"})

	// Positive cards
	DB.Create(&model.Card{Title: "Convincing", Category: model.CardCategoryPositive, EpicID: 1})
	DB.Create(&model.Card{Title: "Fun", Category: model.CardCategoryPositive, EpicID: 2})
	DB.Create(&model.Card{Title: "Good communicator", Category: model.CardCategoryPositive, EpicID: 2})
	DB.Create(&model.Card{Title: "Creative", Category: model.CardCategoryPositive, EpicID: 2})
	DB.Create(&model.Card{Title: "Practical", Category: model.CardCategoryPositive, EpicID: 3})
	DB.Create(&model.Card{Title: "Good listener", Category: model.CardCategoryPositive, EpicID: 2})
	DB.Create(&model.Card{Title: "Straight forward", Category: model.CardCategoryPositive, EpicID: 4})
	DB.Create(&model.Card{Title: "Goal-oriented", Category: model.CardCategoryPositive, EpicID: 5})
	DB.Create(&model.Card{Title: "Tactful", Category: model.CardCategoryPositive, EpicID: 1})
	DB.Create(&model.Card{Title: "Sympathetic", Category: model.CardCategoryPositive, EpicID: 2})
	DB.Create(&model.Card{Title: "Helpful", Category: model.CardCategoryPositive, EpicID: 2})
	DB.Create(&model.Card{Title: "Ambitious", Category: model.CardCategoryPositive, EpicID: 5})
	DB.Create(&model.Card{Title: "Inspiring", Category: model.CardCategoryPositive, EpicID: 1})
	DB.Create(&model.Card{Title: "Optimistic", Category: model.CardCategoryPositive, EpicID: 2})
	DB.Create(&model.Card{Title: "Courageous", Category: model.CardCategoryPositive, EpicID: 5})
	DB.Create(&model.Card{Title: "Determined", Category: model.CardCategoryPositive, EpicID: 5})
	DB.Create(&model.Card{Title: "Emphatic", Category: model.CardCategoryPositive, EpicID: 2})
	DB.Create(&model.Card{Title: "Relaxed", Category: model.CardCategoryPositive, EpicID: 4})
	DB.Create(&model.Card{Title: "Patient", Category: model.CardCategoryPositive, EpicID: 4})
	DB.Create(&model.Card{Title: "Well organized", Category: model.CardCategoryPositive, EpicID: 3})
	DB.Create(&model.Card{Title: "Reliable", Category: model.CardCategoryPositive, EpicID: 3})
	DB.Create(&model.Card{Title: "Independent", Category: model.CardCategoryPositive, EpicID: 3})
	DB.Create(&model.Card{Title: "Flexible", Category: model.CardCategoryPositive, EpicID: 5})
	DB.Create(&model.Card{Title: "Consistent", Category: model.CardCategoryPositive, EpicID: 5})
	DB.Create(&model.Card{Title: "Coaching", Category: model.CardCategoryPositive, EpicID: 1})
	DB.Create(&model.Card{Title: "Hard working", Category: model.CardCategoryPositive, EpicID: 3})
	DB.Create(&model.Card{Title: "Enthousiastic", Category: model.CardCategoryPositive, EpicID: 1})
	DB.Create(&model.Card{Title: "Thoughtful", Category: model.CardCategoryPositive, EpicID: 2})
	DB.Create(&model.Card{Title: "Team player", Category: model.CardCategoryPositive, EpicID: 2})
	DB.Create(&model.Card{Title: "Motivated", Category: model.CardCategoryPositive, EpicID: 1})
	DB.Create(&model.Card{Title: "Knowledgeable", Category: model.CardCategoryPositive, EpicID: 5})
	DB.Create(&model.Card{Title: "Versatile", Category: model.CardCategoryPositive, EpicID: 4})

	// Negative cards
	DB.Create(&model.Card{Title: "Disrespectful", Category: model.CardCategoryNegative, EpicID: 2})
	DB.Create(&model.Card{Title: "Irresponsible", Category: model.CardCategoryNegative, EpicID: 3})
	DB.Create(&model.Card{Title: "Bad communicator", Category: model.CardCategoryNegative, EpicID: 2})
	DB.Create(&model.Card{Title: "Bad listener", Category: model.CardCategoryNegative, EpicID: 2})
	DB.Create(&model.Card{Title: "Dishonest", Category: model.CardCategoryNegative, EpicID: 4})
	DB.Create(&model.Card{Title: "Insecure", Category: model.CardCategoryNegative, EpicID: 1})
	DB.Create(&model.Card{Title: "Tactless", Category: model.CardCategoryNegative, EpicID: 1})
	DB.Create(&model.Card{Title: "Destructive", Category: model.CardCategoryNegative, EpicID: 2})
	DB.Create(&model.Card{Title: "Unmotivated", Category: model.CardCategoryNegative, EpicID: 1})
	DB.Create(&model.Card{Title: "Uninspiring", Category: model.CardCategoryNegative, EpicID: 1})
	DB.Create(&model.Card{Title: "Pessimistic", Category: model.CardCategoryNegative, EpicID: 2})
	DB.Create(&model.Card{Title: "Sloppy", Category: model.CardCategoryNegative, EpicID: 3})
	DB.Create(&model.Card{Title: "Over-sensitive", Category: model.CardCategoryNegative, EpicID: 4})
	DB.Create(&model.Card{Title: "Indecisive", Category: model.CardCategoryNegative, EpicID: 3})
	DB.Create(&model.Card{Title: "Stressed out", Category: model.CardCategoryNegative, EpicID: 4})
	DB.Create(&model.Card{Title: "Intolerant", Category: model.CardCategoryNegative, EpicID: 2})
	DB.Create(&model.Card{Title: "Impatient", Category: model.CardCategoryNegative, EpicID: 4})
	DB.Create(&model.Card{Title: "Chaotic", Category: model.CardCategoryNegative, EpicID: 3})
	DB.Create(&model.Card{Title: "Vague", Category: model.CardCategoryNegative, EpicID: 1})
	DB.Create(&model.Card{Title: "Dependent", Category: model.CardCategoryNegative, EpicID: 3})
	DB.Create(&model.Card{Title: "Rigid", Category: model.CardCategoryNegative, EpicID: 4})
	DB.Create(&model.Card{Title: "Inconsistent", Category: model.CardCategoryNegative, EpicID: 4})
	DB.Create(&model.Card{Title: "Narrow-minded", Category: model.CardCategoryNegative, EpicID: 5})
	DB.Create(&model.Card{Title: "Reluctant", Category: model.CardCategoryNegative, EpicID: 3})
	DB.Create(&model.Card{Title: "Cynical", Category: model.CardCategoryNegative, EpicID: 5})

}
