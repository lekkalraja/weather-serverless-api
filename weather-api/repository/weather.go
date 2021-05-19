package repository

import (
	"io/ioutil"
	"net/http"
)

func GetWeatherResponse(url string) (string, error) {

	req, _ := http.NewRequest("GET", url, nil)
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
