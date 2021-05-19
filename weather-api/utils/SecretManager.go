package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func GetToken(tokenPath string, sess *session.Session) string {

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
