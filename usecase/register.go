package usecase

import (
	"encoding/json"
	"gin-cognito/infra"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func MakeNewUser(c *gin.Context) error {
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
		return err
	}

	return nil
}

type UserPassword struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func EmailTest(c *gin.Context) error {
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
		return err
	}

	return nil
}
