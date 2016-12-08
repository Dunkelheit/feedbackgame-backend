package router

import (
	"net/http"

	"github.com/Dunkelheit/feedbackapp/controller"
	"github.com/Dunkelheit/feedbackapp/model"
	"github.com/Dunkelheit/feedbackgame-backend/router/middleware"
	"gopkg.in/gin-gonic/gin.v1"
)

func ping(c *gin.Context) {
	card := model.Card{Title: "Hello"}
	c.JSON(http.StatusOK, card)
}

func loadRoutes(router *gin.Engine) {
	routes := router.Group("/api")
	routes.POST("/login", controller.Login)

	myRoutes := routes.Group("/my")
	{
		myRoutes.GET("/reviews", middleware.RolesRequired("user"), controller.MyReviews)
	}

	userRoutes := routes.Group("/users")
	{
		userRoutes.GET("", middleware.RolesRequired("user"), controller.AllUsers)
		userRoutes.GET("/:userId", middleware.RolesRequired("user"), controller.UserByID)
		userRoutes.PUT("/:userId", middleware.RolesRequired("user", "admin"), ping)
		userRoutes.DELETE("/:userId", middleware.RolesRequired("user", "admin"), controller.DeleteUser)
	}
	cardRoutes := routes.Group("/cards")
	{
		cardRoutes.GET("", middleware.RolesRequired("user"), controller.AllCards)
		cardRoutes.POST("", middleware.RolesRequired("user", "admin"), controller.CreateCard)
		cardRoutes.PUT("/:cardId", middleware.RolesRequired("user", "admin"), controller.UpdateCard)
		cardRoutes.DELETE("/:cardId", middleware.RolesRequired("user", "admin"), controller.DeleteCard)
	}
	reviewRoutes := routes.Group("/reviews")
	{
		reviewRoutes.GET("", middleware.RolesRequired("user"), controller.AllReviews)
		reviewRoutes.POST("", middleware.RolesRequired("user", "admin"), controller.CreateReview)
		reviewRoutes.GET("/:reviewId", middleware.RolesRequired("user"), ping)
		reviewRoutes.PUT("/:reviewId", middleware.RolesRequired("user"), ping)
		reviewRoutes.DELETE("/:reviewId", middleware.RolesRequired("user", "admin"), ping)
	}
}

// NewRouter makes the router
func NewRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.ParseAuthenticationToken())

	loadRoutes(router)

	router.StaticFile("/", "/Users/arturo.martinez/Projects/feedbackgame-frontend/dist/index.html")
	router.Static("/static", "/Users/arturo.martinez/Projects/feedbackgame-frontend/dist/static/")

	return router
}
