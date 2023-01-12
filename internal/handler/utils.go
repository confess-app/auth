package handler

import (
	"auth/model"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/valyala/fastjson"
)

func CreateResponse(msg string, code int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       msg,
		StatusCode: code,
	}, nil
}

// easy encode
func GenerateToken(item *model.User) (string, error) {
	now := time.Now()
	ttlStr := os.Getenv("TTL_TOKEN")
	ttl, _ := strconv.Atoi(ttlStr)
	expireTime := now.Add(time.Duration(ttl) * time.Second)
	salt := expireTime.Unix()
	b, err := json.Marshal(item)
	if err != nil {
		return "", err
	}
	sEnc := base64.StdEncoding.EncodeToString([]byte(b))
	encode1 := fmt.Sprintf("%d%s", salt, sEnc)
	encode2 := base64.StdEncoding.EncodeToString([]byte(encode1))
	token := encode2 + RandStringRunes(10)
	return token, nil
}

func DecodeTokenToUserModel(token string) (*model.User, error) {
	decode1 := token[0 : len(token)-10]
	decode2, err := base64.StdEncoding.DecodeString(decode1)
	if err != nil {
		return nil, err
	}
	removeSalt := string(decode2)[10:]
	unixStr := string(decode2)[0:10]
	i, err := strconv.ParseInt(unixStr, 10, 64)
	if err != nil {
		return nil, errors.New("wrong token")
	}
	tm := time.Unix(i, 0)
	if time.Now().After(tm) {
		return nil, errors.New("token expired")
	}
	finalStr, err := base64.StdEncoding.DecodeString(string(removeSalt))
	if err != nil {
		return nil, err
	}
	data := &model.User{}
	err = json.Unmarshal([]byte(finalStr), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func ParseData(body string, model interface{}) error {
	err := fastjson.Validate(body)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(body), &model)
	if err != nil {
		return err
	}
	return nil
}
