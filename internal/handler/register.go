package handler

import (
	"encoding/json"
	"net/http"

	"github.com/valyala/fastjson"

	"github.com/aws/aws-lambda-go/events"
)

func Register(body string) (events.APIGatewayProxyResponse, error) {
	err := fastjson.Validate(body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "error: invalid json payload",
			StatusCode: http.StatusInternalServerError,
		}, nil
	}
	data := RegisterData{}
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "error: invalid json payload",
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       data.UserName + "|" + data.Password + "|" + data.Email,
		StatusCode: http.StatusOK,
	}, nil
}

type RegisterData struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
