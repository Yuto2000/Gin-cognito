package usecase

import (
	"encoding/json"
	"gin-cognito/infra"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type UserLogin struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Login(c *gin.Context) (string, error) {
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
	return token, err
}
