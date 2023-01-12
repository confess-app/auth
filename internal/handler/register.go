package handler

import (
	"auth/model"
	"auth/service/mysql"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
	"github.com/valyala/fastjson"
)

func Register(body string) (events.APIGatewayProxyResponse, error) {
	fmt.Println("process register")
	data, err := ParseRegisterData(body)
	if err != nil {
		fmt.Println(err.Error())
		return CreateResponse(err.Error(), http.StatusInternalServerError)
	}
	msg, code, ok := ValidateData(data)
	if !ok {
		fmt.Println("validate data failed: " + msg)
		return CreateResponse(msg, code)
	}

	newUserID, err := uuid.NewRandom()
	if err != nil {
		return CreateResponse(err.Error(), http.StatusInternalServerError)
	}
	hasher := sha1.New()
	hasher.Write([]byte(data.Password))
	pwdSHA := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
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
	return events.APIGatewayProxyResponse{
		Body:       token,
		StatusCode: http.StatusOK,
	}, nil
}

func ParseRegisterData(body string) (*RegisterData, error) {
	err := fastjson.Validate(body)
	if err != nil {
		return nil, err
	}
	data := RegisterData{}
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
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
