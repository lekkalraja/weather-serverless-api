package repository

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type Dynamo struct {
	Client    dynamodbiface.DynamoDBAPI
	TableName string
}

type Item struct {
	Country   string
	Timestamp string
	Data      string
}

func (dynamo Dynamo) CreateItem(item Item) error {

	record, err := dynamodbattribute.MarshalMap(item)

	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      record,
		TableName: aws.String(dynamo.TableName),
	}

	_, err = dynamo.Client.PutItem(input)

	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
		return err
	}

	log.Println("Successfully added '" + item.Country)
	return nil
}
