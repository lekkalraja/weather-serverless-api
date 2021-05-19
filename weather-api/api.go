package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/ssm"
)

var url = ""
var dynamoClient *dynamodb.DynamoDB

type Item struct {
	Country   string
	Timestamp string
	Data      string
}

func init() {

	sess, err := session.NewSession()

	if err != nil {
		panic(err)
	}

	url = fmt.Sprintf("http://%s/%s?q=<country>&appid=%s",
		os.Getenv("HOST"), os.Getenv("ENDPOINT"), getToken(os.Getenv("TOKEN"), sess))

	dynamoClient = dynamodb.New(sess)
}

func getToken(tokenPath string, sess *session.Session) string {

	paramStore := ssm.New(sess, aws.NewConfig())

	param, err := paramStore.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(tokenPath),
		WithDecryption: aws.Bool(true),
	})

	if err != nil {
		panic(err)
	}

	return *param.Parameter.Value
}

func createItem(item Item) {

	av, err := dynamodbattribute.MarshalMap(item)

	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(os.Getenv("TABLE_NAME")),
	}

	_, err = dynamoClient.PutItem(input)

	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}

	fmt.Println("Successfully added '" + item.Country)
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

	item := Item{
		Country:   country,
		Timestamp: time.Now().String(),
		Data:      string(body),
	}

	createItem(item)

	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
