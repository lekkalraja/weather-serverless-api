package main

import (
	"fmt"
	"os"
	"strings"
	"weather-api/repository"
	"weather-api/utils"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var url = ""
var dynamoClient *dynamodb.DynamoDB

func init() {

	sess, err := session.NewSession()

	if err != nil {
		panic(err)
	}

	url = fmt.Sprintf("http://%s/%s?q=<country>&appid=%s",
		os.Getenv("HOST"), os.Getenv("ENDPOINT"), utils.GetToken(os.Getenv("TOKEN"), sess))

	dynamoClient = dynamodb.New(sess)
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	country := request.PathParameters["country"]
	formattedURL := strings.ReplaceAll(url, "<country>", country)
	weather, err := repository.GetWeatherResponse(formattedURL)

	if err != nil {
		return events.APIGatewayProxyResponse{}, nil
	}

	repository.CreateItem(country, weather, dynamoClient)

	return events.APIGatewayProxyResponse{
		Body:       weather,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
