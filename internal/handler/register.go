package handler

import (
	"auth/model"
	"auth/service/mysql"
	"fmt"
	"net/http"
	"net/mail"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
)

func Register(body string) (events.APIGatewayProxyResponse, error) {
	data := RegisterData{}
	err := ParseData(body, &data)
	if err != nil {
		fmt.Println(err.Error())
		return CreateResponse(err.Error(), http.StatusInternalServerError)
	}
	msg, code, ok := ValidateData(&data)
	if !ok {
		fmt.Println("validate data failed: " + msg)
		return CreateResponse(msg, code)
	}

	newUserID, err := uuid.NewRandom()
	if err != nil {
		return CreateResponse(err.Error(), http.StatusInternalServerError)
	}
	pwdSHA := SHAString(data.Password)
	userItem := model.User{
		UserID:   newUserID.String(),
		Username: data.UserName,
		Email:    data.Email,
		Password: pwdSHA,
		Verified: false,
	}
	// TODO: save account
	err = mysql.Save(&userItem)
	if err != nil {
		return CreateResponse(err.Error(), http.StatusInternalServerError)
	}
	token, err := GenerateToken(&userItem)
	if err != nil {
		return CreateResponse(err.Error(), http.StatusInternalServerError)
	}
	return CreateResponse(token, http.StatusOK)
}

func ValidateData(data *RegisterData) (string, int, bool) {
	//validate email
	_, err := mail.ParseAddress(data.Email)
	if err != nil {
		return "error: invalid email format", http.StatusInternalServerError, false
	}
	//check username exists
	user, err := mysql.QueryUserByUsername(data.UserName)
	if err != nil {
		return err.Error(), http.StatusInternalServerError, false
	}
	if user.ID != 0 {
		return "error: username is existed", http.StatusConflict, false
	}
	//check email exists
	user, err = mysql.QueryUserByEmail(data.Email)
	if err != nil {
		return err.Error(), http.StatusInternalServerError, false
	}
	if user.ID != 0 {
		return "error: email is existed", http.StatusConflict, false
	}

	return "", 0, true
}

type RegisterData struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
