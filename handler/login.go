package handler

import (
	"fmt"
	"gin-cognito/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	token, err := usecase.Login(c)
	fmt.Println("token : ", token)

	if err != nil {
		fmt.Println("login error : ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500",
			"messase": "internal server error",
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "make new user ok",
			"token": token,
		})
	}
}
