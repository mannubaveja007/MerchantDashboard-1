package config

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
)

func InitDB() *dynamodb.DynamoDB {
    sess := session.Must(session.NewSession(&aws.Config{
        Region: aws.String("us-east-2"),
    }))
    return dynamodb.New(sess)
}
