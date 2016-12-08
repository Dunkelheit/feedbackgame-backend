package main

import (
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"fmt"

	"github.com/Dunkelheit/feedbackapp/controller"
	"github.com/Dunkelheit/feedbackapp/database"
	"github.com/Dunkelheit/feedbackapp/model"
	"github.com/Dunkelheit/feedbackapp/util"
	"gopkg.in/gin-gonic/gin.v1"
)

func ping(c *gin.Context) {
	card := model.Card{Title: "Hello"}
	c.JSON(http.StatusOK, card)
}

// AuthRequired is the authentication middleware
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("x-auth-token")
		if len(tokenString) == 0 {
			c.JSON(http.StatusUnauthorized, "Authentication required")
			c.Abort()
			return
		}
		username, mail, err := util.DecodeToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "Invalid authentication")
			c.Abort()
			return
		}
		c.Set("username", username)
		fmt.Println("Setting username in middleware")
		fmt.Println(username)
		c.Set("mail", mail)
		c.Next()
	}
}

func main() {
	router := gin.Default()

	apiRoutes := router.Group("/api")
	apiRoutes.POST("/login", controller.Login)

	myRoutes := apiRoutes.Group("/my").Use(AuthRequired())
	{
		myRoutes.GET("/reviews", controller.MyReviews).Use(AuthRequired())
	}

	adminRoutes := apiRoutes.Group("/admin")
	userRoutes := adminRoutes.Group("/users").Use(AuthRequired())
	{
		userRoutes.GET("", controller.AllUsers)
		userRoutes.GET("/:userId", controller.UserByID)
		userRoutes.PUT("/:userId", ping)
		userRoutes.DELETE("/:userId", controller.DeleteUser)
	}
	cardRoutes := adminRoutes.Group("/cards").Use(AuthRequired())
	{
		cardRoutes.GET("", controller.AllCards)
		cardRoutes.POST("", controller.CreateCard)
		cardRoutes.PUT("/:cardId", controller.UpdateCard)
		cardRoutes.DELETE("/:cardId", controller.DeleteCard)
	}
	reviewRoutes := adminRoutes.Group("/reviews").Use(AuthRequired())
	{
		reviewRoutes.GET("", controller.AllReviews)
		reviewRoutes.POST("", controller.CreateReview)
		reviewRoutes.GET("/:reviewId", ping)
		reviewRoutes.PUT("/:reviewId", ping)
		reviewRoutes.DELETE("/:reviewId", ping)
	}

	router.StaticFile("/", "./web/dist/index.html")
	router.Static("/static", "./web/dist/static/")

	router.Run() // listen and server on 0.0.0.0:8080
}

func init() {
	database.OpenDB()
}
