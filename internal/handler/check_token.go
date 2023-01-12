package handler

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func CheckToken(body string) (events.APIGatewayProxyResponse, error) {
	data := VerifyData{}
	err := ParseData(body, &data)
	if err != nil {
		fmt.Println(err.Error())
		return CreateResponse(err.Error(), http.StatusInternalServerError)
	}
	_, err = DecodeTokenToUserModel(data.Token)
	if err != nil {
		fmt.Println(err.Error())
		return CreateResponse(err.Error(), http.StatusNotFound)
	}
	fmt.Println("check token success")
	return CreateResponse("success", http.StatusOK)
}

type VerifyData struct {
	Token string `json:"token"`
}
