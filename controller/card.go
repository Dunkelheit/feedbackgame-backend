package controller

import (
	"net/http"

	"github.com/Dunkelheit/feedbackapp/database"
	"github.com/Dunkelheit/feedbackapp/model"
	"github.com/Dunkelheit/feedbackapp/util"
	"gopkg.in/gin-gonic/gin.v1"
)

// AllCards retrieves all the available cards
func AllCards(c *gin.Context) {
	var cards []model.Card
	database.DB.Order("title asc").Find(&cards)
	c.JSON(http.StatusOK, cards)
}

// CreateCard creates a new card
func CreateCard(c *gin.Context) {
	in := &model.Card{}
	err := c.Bind(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	card := &model.Card{
		Title:    in.Title,
		Category: model.CardCategory(in.Category),
	}
	database.DB.Create(card)

	c.JSON(http.StatusOK, card)
}

// UpdateCard updates a single card
func UpdateCard(c *gin.Context) {
	in := &model.Card{}
	err := c.Bind(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var card model.Card
	if database.DB.First(&card, util.StringToID(c.Param("cardId"))).RecordNotFound() {
		c.JSON(http.StatusNotFound, false)
		return
	}
	card.Title = in.Title
	card.Category = model.CardCategory(in.Category)

	database.DB.Save(&card)
}

// DeleteCard deletes a single card
func DeleteCard(c *gin.Context) {
	var card model.Card
	if database.DB.First(&card, util.StringToID(c.Param("cardId"))).RecordNotFound() {
		c.JSON(http.StatusNotFound, false)
		return
	}
	database.DB.Delete(&card)
	c.JSON(http.StatusOK, true)
}
