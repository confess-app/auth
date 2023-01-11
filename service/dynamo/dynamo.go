package dynamo

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var Client *dynamodb.DynamoDB

const (
	tableName = "user"
)

func Init(region string) {
	sess, err := session.NewSession(&aws.Config{
		// Region: aws.String(os.Getenv("COMMON_REGION"))},
		Region: aws.String(region)},
	)
	if err != nil {
		fmt.Printf("got error when create session config: %s\n", err.Error())
		return
	}
	// Create DynamoDB client
	Client = dynamodb.New(sess)
}

func GetItem(key map[string]*dynamodb.AttributeValue) (*dynamodb.GetItemOutput, error) {
	result, err := Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       key,
	})
	return result, err
}

func PutItem(item map[string]*dynamodb.AttributeValue) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      item,
	}
	_, err := Client.PutItem(input)
	return err
}
