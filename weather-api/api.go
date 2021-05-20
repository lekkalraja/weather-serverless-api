package main

import (
	"weather-api/repository"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	country := request.PathParameters["country"]
	weather, err := repository.GetWeatherResponse(country)

	if err != nil {
		return events.APIGatewayProxyResponse{}, nil
	}

	repository.CreateItem(country, weather)

	return events.APIGatewayProxyResponse{
		Body:       weather,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
