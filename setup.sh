# AWS CLI Command to create encrypted weather token to the secrets maanger parameter store
aws ssm put-parameter --name "/weather/dev/token" --type "SecureString" --value <weather-api-token>