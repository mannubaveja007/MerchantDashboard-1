package models

type Product struct {
    MerchantID string  `json:"merchant_id"`
    ProductID  string  `json:"product_id"`
    Name       string  `json:"name"`
    Price      float64 `json:"price"`
    Quantity   int     `json:"quantity"`
}