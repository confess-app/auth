package main

import (
	"auth/internal/handler"
	"auth/service/mysql"
	"auth/service/redis"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-lambda-go/events"
)

const (
	RegisterPath = "/auth/register"
	VerifyPath   = "/auth/verify"
	LoginPath    = "/auth/login"
	LogoutPath   = "/auth/logout"
)

func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if req.HTTPMethod == http.MethodGet {
		return events.APIGatewayProxyResponse{
			Body:       "error: method get not allowed with this url",
			StatusCode: http.StatusMethodNotAllowed,
		}, nil
	}

	resp := events.APIGatewayProxyResponse{}
	switch req.Path {
	case RegisterPath:
		mysql.Init()
		redis.Init()
		return handler.Register(req.Body)
	case VerifyPath:
	case LoginPath:
	case LogoutPath:
	default:
		return events.APIGatewayProxyResponse{
			Body:       "error: url/api not exists",
			StatusCode: http.StatusNotFound,
		}, nil
	}
	// Response
	return resp, nil
}

func main() {
	lambda.Start(HandleRequest)
	// HandleRequest(events.APIGatewayProxyRequest{
	// 	HTTPMethod: http.MethodPost,
	// 	Path:       RegisterPath,
	// 	Body:       "{ \"username\": \"duong\", \"email\": \"duong@gmail.com\", \"password\": \"password_duong\" }",
	// })
}
