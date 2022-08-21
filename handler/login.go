package handler

import (
	"encoding/json"
	"fmt"
	"gin-cognito/infra"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type UserLogin struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	userPool := infra.NewUserPool()

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic("ioutil readAll error")
	}

	user := UserLogin{}
	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		panic("json unmarshal error")
	}

	// 認証
	token, err := userPool.Authenticate(user.Name, user.Password)
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
		})
	}
}
