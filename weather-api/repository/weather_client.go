package repository

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type WeatherAPI struct {
	Client *http.Client
	URL    string
}

func (api WeatherAPI) GetWeatherResponse(country string) (string, error) {

	req, _ := http.NewRequest("GET", strings.ReplaceAll(api.URL, "<country>", country), nil)
	res, err := api.Client.Do(req)

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
