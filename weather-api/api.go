package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func getUrl(request events.APIGatewayProxyRequest) string {
	country := request.PathParameters["country"]
	host := os.Getenv("HOST")
	path := os.Getenv("ENDPOINT")
	token := os.Getenv("TOKEN")

	return fmt.Sprintf("http://%s/%s?q=%s&appid=%s", host, path, country, token)
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	req, _ := http.NewRequest("GET", getUrl(request), nil)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
