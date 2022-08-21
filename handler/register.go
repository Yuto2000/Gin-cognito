package handler

import (
	"fmt"
	"gin-cognito/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MakeNewUser(c *gin.Context) {
	err := usecase.MakeNewUser(c)

	if err != nil {
		fmt.Println("save error : ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500",
			"messase": "internal server error",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "make new user ok",
		})
	}
}

func EmailTest(c *gin.Context) {
	err := usecase.EmailTest(c)

	if err != nil {
		fmt.Println("email test error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500",
			"message": "internal server error",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "email test ok",
		})
	}
}
