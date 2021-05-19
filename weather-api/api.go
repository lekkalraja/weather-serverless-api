package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

var url = ""

func init() {
	host := os.Getenv("HOST")
	path := os.Getenv("ENDPOINT")
	token := getToken(os.Getenv("TOKEN"))
	fmt.Println(token)
	url = fmt.Sprintf("http://%s/%s?q=<country>&appid=%s", host, path, token)
}

func getToken(tokenPath string) string {

	paramStore := ssm.New(session.New(), aws.NewConfig())

	param, err := paramStore.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(tokenPath),
		WithDecryption: aws.Bool(true),
	})

	if err != nil {
		panic(err)
	}

	return *param.Parameter.Value
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	country := request.PathParameters["country"]
	req, _ := http.NewRequest("GET", strings.ReplaceAll(url, "<country>", country), nil)
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
