package infra

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type UserPool struct {
	client   *cognitoidentityprovider.CognitoIdentityProvider
	PoolId   string
	clientId string
}

type UserInfo struct {
	Uuid    string
	Name    string
	Email   string
	Company string
}

func New(session *session.Session, poolId string, clientId string) *UserPool {
	return &UserPool{
		client:   cognitoidentityprovider.New(session),
		PoolId:   poolId,
		clientId: clientId,
	}
}

func NewUserPool() *UserPool {
	region := os.Getenv("AWS_REGION")
	poolId := os.Getenv("USER_POOL_ID")
	clientId := os.Getenv("APP_CLIENT_ID")
	sess := session.Must(
		session.NewSession(&aws.Config{
			Region: aws.String(region),
		}))
	return New(sess, poolId, clientId)
}

func (p UserPool) Save(name string, email string) error {

	var attrs []*cognitoidentityprovider.AttributeType
	emailAttr := p.userAttribute(cognitoidentityprovider.UsernameAttributeTypeEmail, email)
	attrs = append(attrs, emailAttr)

	input := &cognitoidentityprovider.AdminCreateUserInput{
		UserPoolId:     aws.String(p.PoolId),
		Username:       aws.String(name),
		UserAttributes: attrs,
	}

	_, err := p.client.AdminCreateUser(input)
	if err != nil {
		fmt.Println("adminCreaterUser", err)
		return err
	}

	return nil

}

func (p UserPool) userAttribute(name, value string) *cognitoidentityprovider.AttributeType {
	return &cognitoidentityprovider.AttributeType{
		Name:  aws.String(name),
		Value: aws.String(value),
	}
}

func (p UserPool) Delete(name string) error {

	input := &cognitoidentityprovider.AdminDeleteUserInput{
		UserPoolId: aws.String(p.PoolId),
		Username:   aws.String(name),
	}

	_, err := p.client.AdminDeleteUser(input)
	if err != nil {
		return err
	}

	return nil
}

func (p UserPool) FindOne(name string) (*UserInfo, error) {

	input := &cognitoidentityprovider.AdminGetUserInput{
		UserPoolId: aws.String(p.PoolId),
		Username:   aws.String(name),
	}

	out, err := p.client.AdminGetUser(input)
	if err != nil {
		return nil, err
	}

	userInfo := &UserInfo{
		Name: name,
	}
	for _, a := range out.UserAttributes {
		switch aws.StringValue(a.Name) {
		case "sub":
			userInfo.Uuid = aws.StringValue(a.Value)
		case "email":
			userInfo.Email = aws.StringValue(a.Value)
		case "custom:company":
			userInfo.Company = aws.StringValue(a.Value)
		}
	}

	return userInfo, nil
}

func (p UserPool) SetPassWord(name string, password string) error {

	input := &cognitoidentityprovider.AdminSetUserPasswordInput{
		Password:   aws.String(password),
		Permanent:  aws.Bool(true),
		UserPoolId: aws.String(p.PoolId),
		Username:   aws.String(name),
	}

	_, err := p.client.AdminSetUserPassword(input)
	if err != nil {
		return err
	}

	return nil
}

func (p UserPool) Authenticate(name string, password string) (string, error) {

	input := &cognitoidentityprovider.AdminInitiateAuthInput{
		AuthFlow: aws.String("ADMIN_USER_PASSWORD_AUTH"),
		AuthParameters: map[string]*string{
			"USERNAME": aws.String(name),
			"PASSWORD": aws.String(password),
		},
		ClientId:   aws.String(p.clientId),
		UserPoolId: aws.String(p.PoolId),
	}

	out, err := p.client.AdminInitiateAuth(input)
	if err != nil {
		return "", err
	}

	return aws.StringValue(out.AuthenticationResult.IdToken), nil
}
