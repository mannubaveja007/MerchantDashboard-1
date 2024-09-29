package controllers

import (
	"fmt"
	"merchant-dashboard/models"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
)

// CreateInvoice creates a new invoice
func CreateInvoice(c *gin.Context) {
	var invoice models.Invoice
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save invoice to DynamoDB
	_, err := db.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("InvoicesTable"),
		Item: map[string]*dynamodb.AttributeValue{
			"InvoiceID":  {S: aws.String(invoice.InvoiceID)},
			"MerchantID": {S: aws.String(invoice.MerchantID)},
			"Amount":     {N: aws.String(fmt.Sprintf("%f", invoice.Amount))},
			"Status":     {S: aws.String(invoice.Status)},
		},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create invoice"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Invoice created"})
}

// CheckInvoiceStatus retrieves the status of a specific invoice
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

// UpdateInvoice updates an existing invoice
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
			":amount": {N: aws.String(fmt.Sprintf("%f", invoice.Amount))},
			":status": {S: aws.String(invoice.Status)},
		},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update invoice"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invoice updated"})
}

// DeleteInvoice deletes a specific invoice
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

	c.JSON(http.StatusOK, gin.H{"message": "Invoice deleted"})
}
