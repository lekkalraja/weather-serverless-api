package main

import (
	"os"
	"time"
	repository "weather-api/repository"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var dynamo repository.Dynamo

func init() {
	sess := session.Must(session.NewSession())
	dynamo = repository.Dynamo{
		Client:    dynamodb.New(sess),
		TableName: os.Getenv("TABLE_NAME"),
	}
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	country := request.PathParameters["country"]
	weather, err := repository.GetWeatherResponse(country)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	item := repository.Item{
		Country:   country,
		Timestamp: time.Now().String(),
		Data:      weather,
	}

	dynamoErr := dynamo.CreateItem(item)

	if dynamoErr != nil {
		return events.APIGatewayProxyResponse{}, dynamoErr
	}

	return events.APIGatewayProxyResponse{
		Body:       weather,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
