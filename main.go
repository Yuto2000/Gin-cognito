package main

import (
	// "fmt"
	"gin-cognito/handler"
	"gin-cognito/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("load env error")
	}
	engine := gin.Default()
	ua := ""

	engine.Use(middleware.Middleware)

	testEngine := engine.Group("/test")
	{
		v1 := testEngine.Group("/v1")
		{
			v1.POST("/user", handler.MakeNewUser)
			v1.POST("/email", handler.EmailTest)
			v1.POST("/login", handler.Login)
			// v1.GET("/get", handler.Get)
		}
	}

	// engine.LoadHTMLGlob("templates/*")
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":    "hello world",
			"User-Agent": ua,
		})
		// c.HTML(http.StatusOK, "index.html", gin.H {
		// 	"message": "hello gin",
		// 	"ua": ua,
		// })
	})
	engine.Run(":3000")
}
