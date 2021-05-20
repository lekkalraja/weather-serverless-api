package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
)

type ParamStore struct {
	StoreClient *ssm.SSM
}

func (paramStore *ParamStore) GetToken(tokenPath string) (string, error) {

	param, err := paramStore.StoreClient.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(tokenPath),
		WithDecryption: aws.Bool(true),
	})

	if err != nil {
		return "", err
	}

	return *param.Parameter.Value, nil
}
