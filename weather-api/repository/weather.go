package repository

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"weather-api/utils"

	"github.com/aws/aws-sdk-go/aws/session"
)

var url string

func init() {
	sess, err := session.NewSession()

	if err != nil {
		panic(err)
	}

	url = fmt.Sprintf("http://%s/%s?q=<country>&appid=%s",
		os.Getenv("HOST"), os.Getenv("ENDPOINT"), utils.GetToken(os.Getenv("TOKEN"), sess))

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
