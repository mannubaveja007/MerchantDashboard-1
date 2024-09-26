# Merchant Dashboard API Documentation

## Overview

The Merchant Dashboard is a web-based application that provides merchants with tools to manage their products, subscriptions, payments, and invoices. This API allows merchants to perform CRUD operations on their resources in a secure and efficient manner.

### Features
- **Product Management**: Create, update, retrieve, and delete products.
- **Payment Processing**: Transfer and receive payments.
- **Invoice Management**: Create, check status, update, and delete invoices.
- **Subscription Management**: Manage subscription plans.

## Installation

### Prerequisites
- Go (version 1.16 or later)
- AWS SDK for Go
- DynamoDB set up in AWS
- Gin Gonic framework

### Steps to Install

1. **Clone the Repository**
   ```bash
   git clone <repository-url>
   cd merchant-dashboard
   ```

2. **Install Dependencies**
   ```bash
   go mod tidy
   ```

3. **Set Up AWS Credentials**
   Make sure you have AWS credentials configured on your machine. You can set them up using the AWS CLI:
   ```bash
   aws configure
   ```

4. **Run the Application**
   ```bash
   go run main.go
   ```

5. **Access the API**
   The API will be available at `http://localhost:8080`.

## API Reference

### Authentication

- **Login**
  - **POST** `/auth/login`
  - **Request Body**:
    ```json
    {
      "username": "string",
      "password": "string"
    }
    ```
  - **Response**:
    - `200 OK`: Login successful
    - `400 Bad Request`: Invalid credentials

### Product Management

- **Create Product**
  - **POST** `/products`
  - **Request Body**:
    ```json
    {
      "merchant_id": "string",
      "product_id": "string",
      "name": "string",
      "price": number,
      "quantity": number
    }
    ```
  - **Response**:
    - `201 Created`: Product created
    - `400 Bad Request`: Invalid input

- **Get Products**
  - **GET** `/products`
  - **Query Parameter**: `merchant_id`
  - **Response**:
    - `200 OK`: List of products
    - `500 Internal Server Error`: Could not retrieve products

- **Update Product**
  - **PUT** `/products/:merchantId/:productId`
  - **Request Body**:
    ```json
    {
      "name": "string",
      "price": number,
      "quantity": number
    }
    ```
  - **Response**:
    - `200 OK`: Product updated
    - `400 Bad Request`: Invalid input

- **Delete Product**
  - **DELETE** `/products/:merchantId/:productId`
  - **Response**:
    - `200 OK`: Product deleted
    - `500 Internal Server Error`: Could not delete product

### Payment Processing

- **Transfer Money**
  - **POST** `/payments/transfer`
  - **Request Body**:
    ```json
    {
      "amount": number,
      "bank_account": "string"
    }
    ```
  - **Response**:
    - `200 OK`: Money transferred successfully
    - `400 Bad Request`: Invalid input

- **Receive Money**
  - **POST** `/payments/receive`
  - **Response**:
    - `200 OK`: Money received successfully
    - `400 Bad Request`: Invalid input

### Invoice Management

- **Create Invoice**
  - **POST** `/invoices`
  - **Request Body**:
    ```json
    {
      "invoice_id": "string",
      "merchant_id": "string",
      "amount": number,
      "status": "string"
    }
    ```
  - **Response**:
    - `201 Created`: Invoice created
    - `400 Bad Request`: Invalid input

- **Check Invoice Status**
  - **GET** `/invoices/:invoiceId`
  - **Response**:
    - `200 OK`: Invoice details
    - `404 Not Found`: Invoice not found

- **Update Invoice**
  - **PUT** `/invoices/:invoiceId`
  - **Request Body**:
    ```json
    {
      "amount": number,
      "status": "string"
    }
    ```
  - **Response**:
    - `200 OK`: Invoice updated
    - `400 Bad Request`: Invalid input

- **Delete Invoice**
  - **DELETE** `/invoices/:invoiceId`
  - **Response**:
    - `200 OK`: Invoice deleted
    - `500 Internal Server Error`: Could not delete invoice

### Subscription Management

- **Create Subscription**
  - **POST** `/subscriptions`
  - **Request Body**:
    ```json
    {
      "plan_id": "string",
      "customer_id": "string",
      "price": number
    }
    ```
  - **Response**:
    - `201 Created`: Subscription created
    - `400 Bad Request`: Invalid input

- **Get Subscription**
  - **GET** `/subscriptions/:planId`
  - **Response**:
    - `200 OK`: Subscription details
    - `404 Not Found`: Subscription not found

- **Update Subscription**
  - **PUT** `/subscriptions/:planId`
  - **Request Body**:
    ```json
    {
      "price": number
    }
    ```
  - **Response**:
    - `200 OK`: Subscription updated
    - `400 Bad Request`: Invalid input

- **Delete Subscription**
  - **DELETE** `/subscriptions/:planId`
  - **Response**:
    - `200 OK`: Subscription cancelled
    - `500 Internal Server Error`: Could not delete subscription

## Error Handling

All API responses include relevant HTTP status codes and error messages to help users identify and resolve issues.

## Conclusion

This documentation provides a comprehensive overview of the Merchant Dashboard API. For further assistance or feature requests, please contact me on [[Linkedin](https://www.linkedin.com/in/utkarsh-mahajan3528/)]