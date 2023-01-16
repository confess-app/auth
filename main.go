package main

import (
	"auth/internal/handler"
	"auth/service/mysql"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const (
	RegisterPath = "/auth/register"
	CheckToken   = "/auth/check"
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

	switch req.Path {
	case RegisterPath:
		mysql.Init()
		defer mysql.Close()
		return handler.Register(req.Body)
	case CheckToken:
		return handler.CheckToken(req.Body)
	case LoginPath:
		mysql.Init()
		defer mysql.Close()
		return handler.Login(req.Body)
	case LogoutPath:
		return events.APIGatewayProxyResponse{
			Body:       "success",
			StatusCode: http.StatusOK,
		}, nil
	default:
		return events.APIGatewayProxyResponse{
			Body:       "error: url/api not exists",
			StatusCode: http.StatusNotFound,
		}, nil
	}
}

func main() {
	lambda.Start(HandleRequest)
	// HandleRequest(events.APIGatewayProxyRequest{
	// 	HTTPMethod: http.MethodPost,
	// 	Path:       LoginPath,
	// 	Body:       "{ \"username\": \"duong7\", \"password\": \"password_duong\" }",
	// })
}
