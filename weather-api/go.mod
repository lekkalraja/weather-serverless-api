require github.com/aws/aws-lambda-go v1.23.0

require (
	github.com/aws/aws-sdk-go v1.15.77
	weather-api/repository v0.0.0-00010101000000-000000000000
	weather-api/utils v0.0.0-00010101000000-000000000000
)

replace weather-api/utils => ./utils

replace weather-api/repository => ./repository

module weather-api

go 1.16
