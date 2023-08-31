package middleware

import (
	"Business_Management/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"time"
)

func SessionMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		var session string

		if c.Request.Method == "POST" {
			var request struct {
				Session string `json:"session"`
			}
			if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err})
				c.Abort()
				return
			}
			session = request.Session
		} else if c.Request.Method == "GET" {
			session = c.DefaultQuery("session", "")
		}

		expirationTime := model.Sessions[session].ExpirationTime

		if time.Now().After(expirationTime) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token 无效"})
			c.Abort()
			return
		}
		c.Next()
	}
}
