package controllers

import (
	"fmt"
	"net/http"

	"merchant-dashboard/config" 
	"merchant-dashboard/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/confluentinc/confluent-kafka-go/kafka" 
	"github.com/gin-gonic/gin"
)

func init() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))
	db = dynamodb.New(sess)
}

func CreateInvoice(c *gin.Context) {
	var invoice models.Invoice
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("InvoicesTable"),
		Item: map[string]*dynamodb.AttributeValue{
			"InvoiceID":  {S: aws.String(invoice.InvoiceID)},
			"MerchantID": {S: aws.String(invoice.MerchantID)},
			"Amount":     {N: aws.String(fmt.Sprintf("%.2f", invoice.Amount))},
			"Status":     {S: aws.String(invoice.Status)},
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create invoice"})
		return
	}

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &[]string{"invoice_events"}[0], Partition: -1},
		Value:          []byte(fmt.Sprintf("Created invoice: %s", invoice.InvoiceID)),
	}

	if err := config.KafkaProducer.Produce(msg, nil); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not send message to Kafka"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Invoice created"})
}

func CheckInvoiceStatus(c *gin.Context) {
	invoiceID := c.Param("invoiceId")
	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("InvoicesTable"),
		Key: map[string]*dynamodb.AttributeValue{
			"InvoiceID": {S: aws.String(invoiceID)},
		},
	})
	if err != nil || result.Item == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
		return
	}
	c.JSON(http.StatusOK, result.Item)
}

func UpdateInvoice(c *gin.Context) {
	var invoice models.Invoice
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.UpdateItem(&dynamodb.UpdateItemInput{
		TableName: aws.String("InvoicesTable"),
		Key: map[string]*dynamodb.AttributeValue{
			"InvoiceID": {S: aws.String(invoice.InvoiceID)},
		},
		UpdateExpression: aws.String("set Amount = :amount, Status = :status"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":amount": {N: aws.String(fmt.Sprintf("%.2f", invoice.Amount))},
			":status": {S: aws.String(invoice.Status)},
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update invoice"})
		return
	}

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &[]string{"invoice_events"}[0], Partition: -1}, 
		Value:          []byte(fmt.Sprintf("Updated invoice: %s", invoice.InvoiceID)),
	}

	if err := config.KafkaProducer.Produce(msg, nil); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not send message to Kafka"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invoice updated"})
}

func DeleteInvoice(c *gin.Context) {
	invoiceID := c.Param("invoiceId")

	_, err := db.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String("InvoicesTable"),
		Key: map[string]*dynamodb.AttributeValue{
			"InvoiceID": {S: aws.String(invoiceID)},
		},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete invoice"})
		return
	}

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &[]string{"invoice_events"}[0], Partition: -1}, 
		Value:          []byte(fmt.Sprintf("Deleted invoice: %s", invoiceID)),
	}

	if err := config.KafkaProducer.Produce(msg, nil); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not send message to Kafka"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invoice deleted"})
}
