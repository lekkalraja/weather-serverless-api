# AWS CLI Command to create encrypted weather token to the secrets maanger parameter store
aws ssm put-parameter --name "/weather/dev/token" --type "SecureString" --value <weather-api-token>


# commands to create, copy(web-app) and make s3 bucket as static website hosting
aws s3 mb s3://weather-search-tool
aws s3 cp index.html s3://weather-search-tool
aws s3 website s3://weather-search-tool/ --index-document index.html
aws s3api put-object-acl --bucket weather-search-tool --key index.html --acl public-read

#Commands to delete s3 bucket
aws s3 rm s3://weather-search-tool --recursive
aws s3 rb s3://weather-search-tool
