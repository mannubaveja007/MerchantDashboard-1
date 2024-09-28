package controllers

import (
	"bytes"
	"encoding/json"
	"merchant-dashboard/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Helper function to make API requests to PayME
func callPayMEAPI(method, endpoint string, token string, body interface{}) (*http.Response, error) {
	client := &http.Client{}
	var reqBody []byte
	var err error

	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, "https://api.paymefin.tech/api"+endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	return client.Do(req)
}

func TransferMoney(c *gin.Context) {
	var transferData models.TransferRequest
	if err := c.ShouldBindJSON(&transferData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := c.GetHeader("Authorization") // Extract token from request header
	response, err := callPayMEAPI("POST", "/BLOCKCHAIN/TRANSFER/"+transferData.BankAccount+"/"+strconv.FormatFloat(transferData.Amount, 'f', -1, 64), token, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call PayME API"})
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		var errorResponse map[string]interface{}
		json.NewDecoder(response.Body).Decode(&errorResponse)
		c.JSON(response.StatusCode, errorResponse)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Money transferred successfully"})
}


//return their number and qrcode that contains number
func ReceiveMoney(c *gin.Context) {
	var receiveData models.ReceiveRequest
	if err := c.ShouldBindJSON(&receiveData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := c.GetHeader("Authorization") // Extract token from request header
	// Implement the logic to receive money using the PayME API
	// Assuming there's an endpoint for receiving money
	// Adjust 

	// Example:
	response, err := callPayMEAPI("POST", "/BLOCKCHAIN/RECEIVE", token, receiveData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call PayME API"})
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		var errorResponse map[string]interface{}
		json.NewDecoder(response.Body).Decode(&errorResponse)
		c.JSON(response.StatusCode, errorResponse)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Money received successfully"})
}


