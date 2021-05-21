package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	repository "weather-api/repository"
	utils "weather-api/utils"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/ssm"
)

var (
	dynamo     repository.Dynamo
	weatherAPI repository.WeatherAPI
)

func init() {
	sess := session.Must(session.NewSession())
	dynamo = repository.Dynamo{
		Client:    dynamodb.New(sess),
		TableName: os.Getenv("TABLE_NAME"),
	}

	paramStore := utils.ParamStore{
		StoreClient: ssm.New(sess, aws.NewConfig()),
	}

	token, err := paramStore.GetToken(os.Getenv("TOKEN"))

	log.Println("Got Token : ", token)
	if err != nil {
		log.Fatalf("Failed to fetch Tokean : %s", err)
	}

	url := fmt.Sprintf("http://%s/%s?q=<country>&appid=%s", os.Getenv("HOST"), os.Getenv("ENDPOINT"), token)
	weatherAPI = repository.WeatherAPI{
		Client: http.DefaultClient,
		URL:    url,
	}
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	country := request.PathParameters["country"]
	weather, err := weatherAPI.GetWeatherResponse(country)

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
