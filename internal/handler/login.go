package handler

import (
	"auth/service/mysql"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func Login(body string) (events.APIGatewayProxyResponse, error) {
	data := LoginData{}
	err := ParseData(body, &data)
	if err != nil {
		fmt.Println(err.Error())
		return CreateResponse(err.Error(), http.StatusInternalServerError)
	}
	user, err := mysql.QueryUserByUsername(data.UserName)
	if err != nil {
		fmt.Println(err.Error())
		return CreateResponse(err.Error(), http.StatusInternalServerError)
	}
	if user.ID == 0 {
		return CreateResponse("error: user not exist", http.StatusInternalServerError)
	}
	pwdSHA := SHAString(data.Password)
	if pwdSHA != user.Password {
		fmt.Println(data.Password)
		fmt.Println(user.Password)
		return CreateResponse("error: password not correct", http.StatusUnauthorized)
	}
	token, err := GenerateToken(user)
	if err != nil {
		fmt.Println(err.Error())
		return CreateResponse(err.Error(), http.StatusInternalServerError)
	}
	return events.APIGatewayProxyResponse{
		Body:       token,
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Set-Cookie": fmt.Sprintf("token=%s", token),
		},
	}, nil
}

type LoginData struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
