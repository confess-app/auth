package handler

import (
	"github.com/aws/aws-lambda-go/events"
)

func CreateResponse(msg string, code int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       msg,
		StatusCode: code,
	}, nil
}
