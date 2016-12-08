package controller

import (
	"fmt"
	"net/http"

	"github.com/Dunkelheit/feedbackapp/database"
	"github.com/Dunkelheit/feedbackapp/model"
	"gopkg.in/gin-gonic/gin.v1"
)

// AllReviews retrieves all the available reviews
func AllReviews(c *gin.Context) {
	var reviews []model.Review
	/*
		(/Users/arturo.martinez/Projects/go/src/github.com/Dunkelheit/feedbackapp/controller/review.go:24)
		[2016-11-28 11:46:02]  [0.99ms]  SELECT * FROM "reviews"  WHERE "reviews".deleted_at IS NULL

		(/Users/arturo.martinez/Projects/go/src/github.com/Dunkelheit/feedbackapp/controller/review.go:24)
		[2016-11-28 11:46:02]  [2.81ms]  SELECT * FROM "users"  WHERE "users".deleted_at IS NULL AND (("id" IN ('5')))

		(/Users/arturo.martinez/Projects/go/src/github.com/Dunkelheit/feedbackapp/controller/review.go:24)
		[2016-11-28 11:46:02]  [0.98ms]  SELECT * FROM "users"  WHERE "users".deleted_at IS NULL AND (("id" IN ('6')))

		(/Users/arturo.martinez/Projects/go/src/github.com/Dunkelheit/feedbackapp/controller/review.go:24)
		[2016-11-28 11:46:02]  [0.91ms]  SELECT * FROM "cards" INNER JOIN "review_cards" ON "review_cards"."card_id" = "cards"."id"
		WHERE "cards".deleted_at IS NULL AND (("review_cards"."review_id" IN ('1')))
	*/
	database.DB.Model(&reviews).Preload("Reviewer").Preload("Reviewee").Preload("Cards").Find(&reviews)
	c.JSON(http.StatusOK, reviews)
}

// MyReviews shows only my reviews
func MyReviews(c *gin.Context) {
	userName, what := c.Get("username")
	fmt.Println("Getting username in shenanigans")
	fmt.Println(userName)
	fmt.Println(what)

	var reviews []model.Review

	if err := database.DB.Joins("JOIN users ON users.id = reviews.reviewer_id").
		Where("users.username = ?", userName).Preload("Reviewer").
		Preload("Reviewee").Preload("Cards").Find(&reviews).Error; err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, reviews)
}

// CreateReview creates a review
func CreateReview(c *gin.Context) {
	type reviewForm struct {
		ReviewerID uint `json:"reviewerId" binding:"required"`
		RevieweeID uint `json:"revieweeId" binding:"required"`
	}
	in := &reviewForm{}
	err := c.Bind(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	review := &model.Review{
		Remark:     "Lorem ipsum",
		Completed:  false,
		ReviewerID: in.ReviewerID,
		RevieweeID: in.RevieweeID,
		Cards:      []model.Card{},
	}
	database.DB.Create(review)

	c.JSON(http.StatusOK, review)
}

// CloseReview closes a review
func CloseReview(c *gin.Context) {

}
