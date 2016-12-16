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

	DB.DropTable(&model.Review{}, &model.Card{}, &model.User{}, &model.Role{}, &model.Epic{})
	DB.AutoMigrate(&model.Card{}, &model.User{}, &model.Review{}, &model.Role{}, &model.Epic{})

	DB.Model(&model.Review{}).AddForeignKey("reviewer_id", "users(id)", "CASCADE", "CASCADE")
	DB.Model(&model.Review{}).AddForeignKey("reviewee_id", "users(id)", "CASCADE", "CASCADE")
	DB.Model(&model.Card{}).AddForeignKey("epic_id", "epic(id)", "CASCADE", "CASCADE")

	DB.Create(&model.Epic{Title: "Group A"})
	DB.Create(&model.Epic{Title: "Group B"})
	DB.Create(&model.Epic{Title: "Group C"})
	DB.Create(&model.Epic{Title: "Group D"})

	// Positive cards
	DB.Create(&model.Card{Title: "Hard-working", Category: model.CardCategoryPositive, EpicID: 1})
	DB.Create(&model.Card{Title: "Respectful", Category: model.CardCategoryPositive, EpicID: 1})
	DB.Create(&model.Card{Title: "Responsible", Category: model.CardCategoryPositive, EpicID: 1})
	DB.Create(&model.Card{Title: "Persuasive", Category: model.CardCategoryPositive, EpicID: 1})
	DB.Create(&model.Card{Title: "Humorous", Category: model.CardCategoryPositive, EpicID: 1})
	DB.Create(&model.Card{Title: "Good communicator", Category: model.CardCategoryPositive, EpicID: 1})
	DB.Create(&model.Card{Title: "Creative", Category: model.CardCategoryPositive, EpicID: 1})
	DB.Create(&model.Card{Title: "Practical", Category: model.CardCategoryPositive, EpicID: 2})
	DB.Create(&model.Card{Title: "Good listener", Category: model.CardCategoryPositive, EpicID: 2})
	DB.Create(&model.Card{Title: "Honest", Category: model.CardCategoryPositive, EpicID: 2})
	DB.Create(&model.Card{Title: "Self-assured", Category: model.CardCategoryPositive, EpicID: 2})
	DB.Create(&model.Card{Title: "Goal-oriented", Category: model.CardCategoryPositive, EpicID: 2})
	DB.Create(&model.Card{Title: "Tactful", Category: model.CardCategoryPositive, EpicID: 2})
	DB.Create(&model.Card{Title: "Sympathetic", Category: model.CardCategoryPositive, EpicID: 2})
	DB.Create(&model.Card{Title: "Helpful", Category: model.CardCategoryPositive, EpicID: 3})
	DB.Create(&model.Card{Title: "Ambitious", Category: model.CardCategoryPositive, EpicID: 3})
	DB.Create(&model.Card{Title: "Inspiring", Category: model.CardCategoryPositive, EpicID: 3})
	DB.Create(&model.Card{Title: "Optimistic", Category: model.CardCategoryPositive, EpicID: 3})
	DB.Create(&model.Card{Title: "Courageous", Category: model.CardCategoryPositive, EpicID: 3})
	DB.Create(&model.Card{Title: "Accurate", Category: model.CardCategoryPositive, EpicID: 3})
	DB.Create(&model.Card{Title: "Determined", Category: model.CardCategoryPositive, EpicID: 3})
	DB.Create(&model.Card{Title: "Sensitive", Category: model.CardCategoryPositive, EpicID: 4})
	DB.Create(&model.Card{Title: "Calm", Category: model.CardCategoryPositive, EpicID: 4})
	DB.Create(&model.Card{Title: "Tolerant", Category: model.CardCategoryPositive, EpicID: 4})
	DB.Create(&model.Card{Title: "Patient", Category: model.CardCategoryPositive, EpicID: 4})
	DB.Create(&model.Card{Title: "Well organized", Category: model.CardCategoryPositive, EpicID: 4})
	DB.Create(&model.Card{Title: "Direct", Category: model.CardCategoryPositive, EpicID: 4})
	DB.Create(&model.Card{Title: "Trustworthy", Category: model.CardCategoryPositive, EpicID: 4})
	DB.Create(&model.Card{Title: "Independent", Category: model.CardCategoryPositive, EpicID: 4})
	DB.Create(&model.Card{Title: "Flexible", Category: model.CardCategoryPositive, EpicID: 4})
	DB.Create(&model.Card{Title: "Consistent", Category: model.CardCategoryPositive, EpicID: 4})

	// Negative cards
	DB.Create(&model.Card{Title: "Disrespectful", Category: model.CardCategoryNegative, EpicID: 1})
	DB.Create(&model.Card{Title: "Irresponsible", Category: model.CardCategoryNegative, EpicID: 1})
	DB.Create(&model.Card{Title: "Unbelievable", Category: model.CardCategoryNegative, EpicID: 1})
	DB.Create(&model.Card{Title: "Serious", Category: model.CardCategoryNegative, EpicID: 1})
	DB.Create(&model.Card{Title: "Bad communicator", Category: model.CardCategoryNegative, EpicID: 1})
	DB.Create(&model.Card{Title: "Unimaginative", Category: model.CardCategoryNegative, EpicID: 1})
	DB.Create(&model.Card{Title: "Theoretical", Category: model.CardCategoryNegative, EpicID: 1})
	DB.Create(&model.Card{Title: "Always talking", Category: model.CardCategoryNegative, EpicID: 2})
	DB.Create(&model.Card{Title: "Dishonest", Category: model.CardCategoryNegative, EpicID: 2})
	DB.Create(&model.Card{Title: "Unsure", Category: model.CardCategoryNegative, EpicID: 2})
	DB.Create(&model.Card{Title: "Lethargic", Category: model.CardCategoryNegative, EpicID: 2})
	DB.Create(&model.Card{Title: "Tactless", Category: model.CardCategoryNegative, EpicID: 2})
	DB.Create(&model.Card{Title: "Unfeeling", Category: model.CardCategoryNegative, EpicID: 2})
	DB.Create(&model.Card{Title: "Hindering", Category: model.CardCategoryNegative, EpicID: 2})
	DB.Create(&model.Card{Title: "Unmotivated", Category: model.CardCategoryNegative, EpicID: 3})
	DB.Create(&model.Card{Title: "Uninspiring", Category: model.CardCategoryNegative, EpicID: 3})
	DB.Create(&model.Card{Title: "Pessimistic", Category: model.CardCategoryNegative, EpicID: 3})
	DB.Create(&model.Card{Title: "Fearful", Category: model.CardCategoryNegative, EpicID: 3})
	DB.Create(&model.Card{Title: "Inaccurate", Category: model.CardCategoryNegative, EpicID: 3})
	DB.Create(&model.Card{Title: "Undetermined", Category: model.CardCategoryNegative, EpicID: 3})
	DB.Create(&model.Card{Title: "Over-sensitive", Category: model.CardCategoryNegative, EpicID: 3})
	DB.Create(&model.Card{Title: "Restless", Category: model.CardCategoryNegative, EpicID: 4})
	DB.Create(&model.Card{Title: "Intolerant", Category: model.CardCategoryNegative, EpicID: 4})
	DB.Create(&model.Card{Title: "Impatient", Category: model.CardCategoryNegative, EpicID: 4})
	DB.Create(&model.Card{Title: "Chaotic", Category: model.CardCategoryNegative, EpicID: 4})
	DB.Create(&model.Card{Title: "Vague", Category: model.CardCategoryNegative, EpicID: 4})
	DB.Create(&model.Card{Title: "Untrustworthy", Category: model.CardCategoryNegative, EpicID: 4})
	DB.Create(&model.Card{Title: "Over-dependent", Category: model.CardCategoryNegative, EpicID: 4})
	DB.Create(&model.Card{Title: "Inflexible", Category: model.CardCategoryNegative, EpicID: 4})
	DB.Create(&model.Card{Title: "Inconsistent", Category: model.CardCategoryNegative, EpicID: 4})
}
