package repository

import (
	"errors"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type MockDynamo struct {
	dynamodbiface.DynamoDBAPI
	Response dynamodb.PutItemOutput
}

func (mockDynamo MockDynamo) PutItem(in *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	if *in.TableName == "ErrorTable" {
		return &dynamodb.PutItemOutput{}, errors.New("Failed To Create Item")
	}
	return &mockDynamo.Response, nil
}

func TestCreateItem(t *testing.T) {
	mock := MockDynamo{
		Response: dynamodb.PutItemOutput{},
	}

	item := Item{
		Country:   "Singapore",
		Timestamp: time.Now().String(),
		Data:      "Response",
	}

	dynamo := Dynamo{
		Client:    mock,
		TableName: "Test_Table",
	}

	t.Run("Should Create Item", func(t *testing.T) {
		err := dynamo.CreateItem(item)
		if err != nil {
			t.Fatal("Should Create Item in the DB")
		}
	})

	t.Run("Should Handle Failures of Create Item", func(t *testing.T) {
		dynamo.TableName = "ErrorTable"
		err := dynamo.CreateItem(item)
		if err == nil {
			t.Fatal("Should Throw err incase of failure of creating item")
		}
	})
}
