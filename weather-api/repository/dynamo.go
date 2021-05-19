package repository

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Item struct {
	Country   string
	Timestamp string
	Data      string
}

func CreateItem(country string, weather string, dynamoClient *dynamodb.DynamoDB) {

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

	fmt.Println("Successfully added '" + item.Country)
}
