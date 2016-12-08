package middleware

import (
	"net/http"

	"github.com/Dunkelheit/feedbackgame-backend/database"
	"github.com/Dunkelheit/feedbackgame-backend/model"
	"github.com/Dunkelheit/feedbackgame-backend/util"
	"gopkg.in/gin-gonic/gin.v1"
)

// ParseAuthenticationToken parses the authentication token, if present
func ParseAuthenticationToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("x-auth-token")
		if len(tokenString) == 0 {
			c.Next()
			return
		}
		username, mail, err := util.DecodeToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "Invalid authentication token")
			c.Abort()
			return
		}
		c.Set("username", username)
		c.Set("mail", mail)
		c.Next()
	}
}

// RolesRequired when you require some roles -- TODO: proper implementation
func RolesRequired(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, _ := c.Get("username")

		switch len(roles) {
		case 1:
			if username != nil {
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, "Not authorized")
				c.Abort()
			}
		case 2:
			var roles []model.Role
			database.DB.Where("role = ? AND username = ?", "admin", username).Find(&roles)
			if len(roles) > 0 {
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, "Not authorized")
				c.Abort()
			}
		}
		c.Next()
	}
}
