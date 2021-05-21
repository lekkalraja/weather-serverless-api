package repository

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type MockHttpClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (c *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	println()
	var json string
	country := req.URL.Query().Get("q")
	if country == "InvalidCountry" {
		return &http.Response{}, errors.New("Something went wrong")
	}
	if country == "Hyderabad" {
		json = `
			{
				"weather": [
					{
						"id": 803,
						"main": "Clouds",
						"description": "broken clouds",
						"icon": "04d"
					}
				]
			}
		`
	} else {
		json = `
			{
				"cod": "404",
				"message": "city not found"
			}
		`
	}
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	return &http.Response{
		StatusCode: 200,
		Body:       r,
	}, nil
}

func TestGetWeatherResponse(t *testing.T) {

	mock := &MockHttpClient{}

	api := WeatherAPI{
		Client: mock,
		URL:    "http://127.0.0.1/weather?q=<country>&appid=test",
	}

	t.Run("Should Return Weather Response For Valid Country", func(t *testing.T) {
		res, err := api.GetWeatherResponse("Hyderabad")
		if !strings.Contains(res, "weather") || err != nil {
			t.Fatal("Should Get Weather Response")
		}
	})

	t.Run("Should Return 404 For Invalid Country", func(t *testing.T) {
		res, err := api.GetWeatherResponse("Vizag")
		if !strings.Contains(res, "404") || err != nil {
			t.Fatal("Should Get Invalid Weather Response")
		}
	})

	t.Run("Should Handle Failure Scenarios", func(t *testing.T) {
		_, err := api.GetWeatherResponse("InvalidCountry")
		if err == nil {
			t.Fatal("Should Throw error")
		}
	})
}
