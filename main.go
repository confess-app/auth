package main

import (
	"auth/internal/handler"
	"auth/service/mysql"
	"auth/service/redis"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"

	_ "github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-lambda-go/events"
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

	resp := events.APIGatewayProxyResponse{}
	switch req.Path {
	case RegisterPath:
		mysql.Init()
		redis.Init()
		return handler.Register(req.Body)
	case CheckToken:
		return handler.CheckToken(req.Body)
	case LoginPath:

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
	// Response
	return resp, nil
}

func main() {
	lambda.Start(HandleRequest)
	// HandleRequest(events.APIGatewayProxyRequest{
	// 	HTTPMethod: http.MethodPost,
	// 	Path:       CheckToken,
	// 	Body:       "{ \"token\": \"MTY3MzQ2OTg5NmV5SkpSQ0k2TkN3aVEzSmxZWFJsWkVGMElqb2lNakF5TXkwd01TMHhNVlF5TURvME1UbzFOaTQ0TURWYUlpd2lWWEJrWVhSbFpFRjBJam9pTWpBeU15MHdNUzB4TVZReU1EbzBNVG8xTmk0NE1EVmFJaXdpUkdWc1pYUmxaRUYwSWpwdWRXeHNMQ0oxYzJWeVgybGtJam9pTW1OaU1EVTNZV1F0TWpReU5DMDBOMll3TFRsbFpEZ3RaakE1TmpGbVl6TTRPV1U0SWl3aWRYTmxjbTVoYldVaU9pSmtkVzl1WnpZaUxDSndZWE56ZDI5eVpDSTZJa0ZGWXpJemMzaHBVbFZVTm0weVNEUTNSVmg0T0hGclUwSk9SVDBpTENKbGJXRnBiQ0k2SW1SMWIyNW5Oa0JuYldGcGJDNWpiMjBpTENKMlpYSnBabWxsWkNJNlptRnNjMlY5juXBhUCBIO\" }",
	// })
}
