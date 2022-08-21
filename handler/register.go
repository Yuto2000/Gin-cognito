package handler

import (
	"encoding/json"
	"fmt"
	"gin-cognito/infra"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func MakeNewUser(c *gin.Context) {
	userPool := infra.NewUserPool()

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic("ioutil readAll error")
	}

	user := User{}
	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		panic("json unmarshal error")
	}

	// user作成
	err = userPool.Save(user.Name, user.Email)

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

type UserPassword struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func EmailTest(c *gin.Context) {
	userPool := infra.NewUserPool()

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic("ioutil readAll error")
	}

	userPassword := UserPassword{}
	err = json.Unmarshal(jsonData, &userPassword)
	if err != nil {
		panic("json unmarshal error")
	}

	err = userPool.SetPassWord(userPassword.Name, userPassword.Password)

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
