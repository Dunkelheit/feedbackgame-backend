package controller

import (
	"net/http"

	"github.com/Dunkelheit/feedbackapp/database"
	"github.com/Dunkelheit/feedbackapp/model"
	"github.com/Dunkelheit/feedbackapp/util"

	"gopkg.in/gin-gonic/gin.v1"
)

// AllUsers retrieves all the available users
func AllUsers(c *gin.Context) {
	var users []model.User
	database.DB.Order("first_name asc").Find(&users)
	c.JSON(http.StatusOK, users)
}

// UserByID gets a single user by its identifier
func UserByID(c *gin.Context) {
	var user model.User
	if database.DB.First(&user, util.StringToID(c.Param("userId"))).RecordNotFound() {
		c.JSON(http.StatusNotFound, "User not found")
		return
	}
	c.JSON(http.StatusOK, user)
}

// DeleteUser deletes a single user
func DeleteUser(c *gin.Context) {
	var user model.User
	if database.DB.First(&user, util.StringToID(c.Param("userId"))).RecordNotFound() {
		c.JSON(http.StatusNotFound, false)
		return
	}
	database.DB.Delete(&user)
	c.JSON(http.StatusOK, true)
}
