package repository

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	utils "weather-api/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

var url string

func init() {
	sess := session.Must(session.NewSession())

	paramStore := utils.ParamStore{
		StoreClient: ssm.New(sess, aws.NewConfig()),
	}

	token, err := paramStore.GetToken(os.Getenv("TOKEN"))

	log.Println("Got Token : ", token)
	if err != nil {
		log.Fatalf("Failed to fetch Tokean : %s", err)
	}

	url = fmt.Sprintf("http://%s/%s?q=<country>&appid=%s", os.Getenv("HOST"), os.Getenv("ENDPOINT"), token)
}

func GetWeatherResponse(country string) (string, error) {

	req, _ := http.NewRequest("GET", strings.ReplaceAll(url, "<country>", country), nil)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, ioErr := ioutil.ReadAll(res.Body)

	if ioErr != nil {
		return "", ioErr
	}

	return string(body), nil
}
