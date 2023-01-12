package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/valyala/fastjson"
)

func CheckToken(body string) (events.APIGatewayProxyResponse, error) {
	err := fastjson.Validate(body)
	if err != nil {
		return CreateResponse(err.Error(), http.StatusBadRequest)
	}
	data := &VerifyData{}
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		return CreateResponse(err.Error(), http.StatusBadRequest)
	}
	_, err = DecodeTokenToUserModel(data.Token)
	if err != nil {
		fmt.Println(err.Error())
		return CreateResponse(err.Error(), http.StatusNotFound)
	}
	return CreateResponse("success", http.StatusOK)
}

type VerifyData struct {
	Token string `json:"token"`
}
