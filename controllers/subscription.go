package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"merchant-dashboard/config" 
	"merchant-dashboard/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/confluentinc/confluent-kafka-go/kafka" 
	"github.com/gin-gonic/gin"
)

var db *dynamodb.DynamoDB

func init() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))
	db = dynamodb.New(sess)
}

func CreateSubscription(c *gin.Context) {
	var subscription models.Subscription
	if err := c.ShouldBindJSON(&subscription); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("SubscriptionsTable"),
		Item: map[string]*dynamodb.AttributeValue{
			"PlanID":     {S: aws.String(subscription.PlanID)},
			"CustomerID": {S: aws.String(subscription.CustomerID)},
			"Price":      {N: aws.String(fmt.Sprintf("%.2f", subscription.Price))},
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create subscription"})
		return
	}

	
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &[]string{"subscription_events"}[0], Partition: 0}, 
		Value:          []byte(fmt.Sprintf("Created subscription for customer: %s", subscription.CustomerID)),
	}

	
	if err := config.KafkaProducer.Produce(msg, nil); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not send message to Kafka"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Subscription created"})
}

func GetSubscription(c *gin.Context) {
	planID := c.Param("plan_id")
	customerID := c.Param("customer_id")
	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("SubscriptionsTable"),
		Key: map[string]*dynamodb.AttributeValue{
			"PlanID":     {S: aws.String(planID)},
			"CustomerID": {S: aws.String(customerID)},
		},
	})
	if err != nil || result.Item == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subscription not found"})
		return
	}
	price, _ := strconv.ParseFloat(*result.Item["Price"].N, 64)
	subscription := models.Subscription{
		PlanID:     *result.Item["PlanID"].S,
		CustomerID: *result.Item["CustomerID"].S,
		Price:      price,
	}
	c.JSON(http.StatusOK, subscription)
}

func UpdateSubscription(c *gin.Context) {
	var subscription models.Subscription
	if err := c.ShouldBindJSON(&subscription); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.UpdateItem(&dynamodb.UpdateItemInput{
		TableName: aws.String("SubscriptionsTable"),
		Key: map[string]*dynamodb.AttributeValue{
			"PlanID":     {S: aws.String(subscription.PlanID)},
			"CustomerID": {S: aws.String(subscription.CustomerID)},
		},
		UpdateExpression: aws.String("set Price = :price"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":price": {N: aws.String(fmt.Sprintf("%.2f", subscription.Price))},
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update subscription"})
		return
	}

	
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &[]string{"subscription_events"}[0], Partition: 0}, 
		Value:          []byte(fmt.Sprintf("Updated subscription for customer: %s", subscription.CustomerID)),
	}

	
	if err := config.KafkaProducer.Produce(msg, nil); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not send message to Kafka"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Subscription updated"})
}

func DeleteSubscription(c *gin.Context) {
	planID := c.Param("plan_id")
	customerID := c.Param("customer_id")
	_, err := db.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String("SubscriptionsTable"),
		Key: map[string]*dynamodb.AttributeValue{
			"PlanID":     {S: aws.String(planID)},
			"CustomerID": {S: aws.String(customerID)},
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete subscription"})
		return
	}

	
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &[]string{"subscription_events"}[0], Partition: 0}, 
		Value:          []byte(fmt.Sprintf("Deleted subscription for customer: %s", customerID)),
	}

	
	if err := config.KafkaProducer.Produce(msg, nil); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not send message to Kafka"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Subscription deleted"})
}
