package controller

import (
	"net/http"

	"github.com/Dunkelheit/feedbackgame-backend/database"
	"github.com/Dunkelheit/feedbackgame-backend/model"
	"gopkg.in/gin-gonic/gin.v1"
)

// AllEpics retrieves all the available epics
func AllEpics(c *gin.Context) {
	var epics []model.Epic
	database.DB.Order("title asc").Find(&epics)
	c.JSON(http.StatusOK, epics)
}
