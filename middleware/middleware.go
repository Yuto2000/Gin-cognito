package middleware

import (
	"github.com/gin-gonic/gin"
)

func Middleware(c *gin.Context) {
	// ua = c.GetHeader("User-Agent")
	c.Next()
}
