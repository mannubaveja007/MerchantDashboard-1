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
		Region: aws.String("us-west-2"),
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
		TableName: aws.String("ProductsTable"),
		Item: map[string]*dynamodb.AttributeValue{
			"MerchantID": {S: aws.String(product.MerchantID)},
			"ProductID":  {S: aws.String(product.ProductID)},
			"Name":       {S: aws.String(product.Name)},
			"Price":      {N: aws.String(fmt.Sprintf("%f", product.Price))},
			"Quantity":   {N: aws.String(fmt.Sprintf("%d", product.Quantity))},
		},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created"})
}

func GetProducts(c *gin.Context) {
	merchantID := c.Query("merchant_id")

	result, err := db.Scan(&dynamodb.ScanInput{
		TableName:        aws.String("ProductsTable"),
		FilterExpression: aws.String("MerchantID = :mid"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":mid": {S: aws.String(merchantID)},
		},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve products"})
		return
	}

	products := make([]models.Product, len(result.Items))
	for i, item := range result.Items {
		price, _ := strconv.ParseFloat(*item["Price"].N, 64)
		quantity, _ := strconv.Atoi(*item["Quantity"].N)
		products[i] = models.Product{
			MerchantID: *item["MerchantID"].S,
			ProductID:  *item["ProductID"].S,
			Name:       *item["Name"].S,
			Price:      price,
			Quantity:   quantity,
		}
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
		TableName: aws.String("ProductsTable"),
		Key: map[string]*dynamodb.AttributeValue{
			"MerchantID": {S: aws.String(product.MerchantID)},
			"ProductID":  {S: aws.String(product.ProductID)},
		},
		UpdateExpression: aws.String("set Name = :name, Price = :price, Quantity = :quantity"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":name":     {S: aws.String(product.Name)},
			":price":    {N: aws.String(fmt.Sprintf("%f", product.Price))},
			":quantity": {N: aws.String(fmt.Sprintf("%d", product.Quantity))},
		},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated"})
}

func DeleteProduct(c *gin.Context) {
	merchantID := c.Param("merchantId")
	productID := c.Param("productId")

	_, err := db.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String("ProductsTable"),
		Key: map[string]*dynamodb.AttributeValue{
			"MerchantID": {S: aws.String(merchantID)},
			"ProductID":  {S: aws.String(productID)},
		},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
