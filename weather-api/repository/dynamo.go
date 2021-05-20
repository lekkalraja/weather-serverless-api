package repository

import (
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var dynamoClient *dynamodb.DynamoDB

func init() {
	sess, err := session.NewSession()

	if err != nil {
		panic(err)
	}

	dynamoClient = dynamodb.New(sess)
}

type Item struct {
	Country   string
	Timestamp string
	Data      string
}

func CreateItem(country string, weather string) {

	item := Item{
		Country:   country,
		Timestamp: time.Now().String(),
		Data:      weather,
	}

	record, err := dynamodbattribute.MarshalMap(item)

	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      record,
		TableName: aws.String(os.Getenv("TABLE_NAME")),
	}

	_, err = dynamoClient.PutItem(input)

	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}

	log.Println("Successfully added '" + item.Country)
}
