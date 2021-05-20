package utils

import (
	"errors"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

type GetParameterMock struct {
	ssmiface.SSMAPI
	Response ssm.GetParameterOutput
}

func (mock GetParameterMock) GetParameter(input *ssm.GetParameterInput) (*ssm.GetParameterOutput, error) {
	if strings.Contains(*input.Name, "Invalid") {
		return &ssm.GetParameterOutput{}, errors.New("Invalid Token")
	}
	output := new(ssm.GetParameterOutput)
	output.Parameter = &ssm.Parameter{Value: aws.String("TestToken1234")}
	return output, nil
}

func TestGetParameter(t *testing.T) {
	expectedToken := "TestToken1234"

	mock := &GetParameterMock{
		Response: ssm.GetParameterOutput{
			Parameter: &ssm.Parameter{
				Value: &expectedToken,
			},
		},
	}

	parameterStore := ParamStore{
		StoreClient: mock,
	}

	t.Run("Should Get Parameter", func(t *testing.T) {
		actualToken, _ := parameterStore.GetToken("/weather/dev/token")
		if actualToken != expectedToken {
			t.Errorf("Expected %q but got %q", expectedToken, actualToken)
		}
	})

	t.Run("Should Get Error For Invalid Token", func(t *testing.T) {
		_, err := parameterStore.GetToken("/weather/dev/Invalidtoken")
		if err == nil {
			t.Errorf("Expected Error Token")
		}

	})
}
