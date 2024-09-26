package utils

import (
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
)

var db *dynamodb.DynamoDB

func InitDB() {
    sess := session.Must(session.NewSession(&aws.Config{
        Region: aws.String("us-west-2"),
    }))
    db = dynamodb.New(sess)
}

func GetDB() *dynamodb.DynamoDB {
    return db
}
