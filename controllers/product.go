package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"merchant-dashboard/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
)


func init() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))
	db = dynamodb.New(sess)
}

// CreateProduct adds a new product to the DynamoDB table
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("products"),
		Item: map[string]*dynamodb.AttributeValue{
			"merchantID": {S: aws.String(product.MerchantID)},
			"productID":  {S: aws.String(product.ProductID)},
			"Name":       {S: aws.String(product.Name)},
			"Price":      {N: aws.String(fmt.Sprintf("%.2f", product.Price))},
			"Quantity":   {N: aws.String(strconv.Itoa(product.Quantity))},
		},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created"})
}

// GetProducts retrieves all products for a given merchant
func GetProducts(c *gin.Context) {
	merchantID := c.Query("merchant_id")
	if merchantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "merchant_id is required"})
		return
	}

	result, err := db.Scan(&dynamodb.ScanInput{
		TableName:        aws.String("products"),
		FilterExpression: aws.String("merchantID = :mid"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":mid": {S: aws.String(merchantID)},
		},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve products", "details": err.Error()})
		return
	}

	products := make([]models.Product, 0) // Initialize an empty slice
	for _, item := range result.Items {
		price, _ := strconv.ParseFloat(*item["Price"].N, 64)
		quantity, _ := strconv.Atoi(*item["Quantity"].N)

		products = append(products, models.Product{
			MerchantID: *item["merchantID"].S,
			ProductID:  *item["productID"].S,
			Name:       *item["Name"].S,
			Price:      price,
			Quantity:   quantity,
		})
	}

	c.JSON(http.StatusOK, products)
}

// UpdateProduct modifies an existing product in the DynamoDB table
func UpdateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.UpdateItem(&dynamodb.UpdateItemInput{
		TableName: aws.String("products"),
		Key: map[string]*dynamodb.AttributeValue{
			"merchantID": {S: aws.String(product.MerchantID)},
			"productID":  {S: aws.String(product.ProductID)},
		},
		UpdateExpression: aws.String("set Name = :name, Price = :price, Quantity = :quantity"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":name":     {S: aws.String(product.Name)},
			":price":    {N: aws.String(fmt.Sprintf("%.2f", product.Price))},
			":quantity": {N: aws.String(strconv.Itoa(product.Quantity))},
		},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not update product: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated"})
}

// DeleteProduct removes a product from the DynamoDB table
func DeleteProduct(c *gin.Context) {
	merchantID := c.Param("merchantId")
	productID := c.Param("productId")

	_, err := db.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String("products"),
		Key: map[string]*dynamodb.AttributeValue{
			"merchantID": {S: aws.String(merchantID)},
			"productID":  {S: aws.String(productID)},
		},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not delete product: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
