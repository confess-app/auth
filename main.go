package main

import (
	auth "auth/internal/handler"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const (
	RegisterPath = "/auth/register"
	VerifyPath   = "/auth/verify"
	LoginPath    = "/auth/login"
	LogoutPath   = "/auth/logout"
)

func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("method=", req.HTTPMethod)
	fmt.Println("path=", req.Path)
	if req.HTTPMethod == http.MethodGet {
		return events.APIGatewayProxyResponse{
			Body:       "error: method get not allowed with this url",
			StatusCode: http.StatusMethodNotAllowed,
		}, nil
	}

	resp := events.APIGatewayProxyResponse{}
	switch req.Path {
	case RegisterPath:
		return auth.Register(req.Body)
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
}
