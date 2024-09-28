package controllers

import (
	"fmt"
	"merchant-dashboard/models"
	"net/http"
	"strconv"

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

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("products"),
		Item: map[string]*dynamodb.AttributeValue{
			"merchantID": {N: aws.String(strconv.Itoa(product.MerchantID))},
			"productID":  {N: aws.String(strconv.Itoa(product.ProductID))},
			"Name":       {S: aws.String(product.Name)},
			"Price":      {N: aws.String(fmt.Sprintf("%.2f", product.Price))}, // Format as float
			"Quantity":   {N: aws.String(strconv.Itoa(product.Quantity))},
		},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created"})
}

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
			":mid": {N: aws.String(merchantID)},
		},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve products", "details": err.Error()})
		return
	}

	products := make([]models.Product, 0) // Use 0 for initial length
	for _, item := range result.Items {
		price, _ := strconv.ParseFloat(*item["Price"].N, 64)
		quantity, _ := strconv.Atoi(*item["Quantity"].N)
		merchantID, _ := strconv.Atoi(*item["merchantID"].N)
		productID, _ := strconv.Atoi(*item["productID"].N)

		products = append(products, models.Product{
			MerchantID: merchantID,
			ProductID:  productID,
			Name:       *item["Name"].S,
			Price:      price,
			Quantity:   quantity,
		})
	}

	c.JSON(http.StatusOK, products)
}

func UpdateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.UpdateItem(&dynamodb.UpdateItemInput{
		TableName: aws.String("products"),
		Key: map[string]*dynamodb.AttributeValue{
			"merchantID": {N: aws.String(strconv.Itoa(product.MerchantID))},
			"productID":  {N: aws.String(strconv.Itoa(product.ProductID))},
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

func DeleteProduct(c *gin.Context) {
	merchantID := c.Param("merchantId")
	productID := c.Param("productId")

	_, err := db.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String("products"),
		Key: map[string]*dynamodb.AttributeValue{
			"merchantID": {N: aws.String(merchantID)},
			"productID":  {N: aws.String(productID)},
		},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not delete product: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
