package models

type Product struct {
    MerchantID int     `json:"merchant_id"`
    ProductID  int     `json:"product_id"`
    Name       string  `json:"name"`
    Price      float64 `json:"price"`   // Change to float64
    Quantity   int     `json:"quantity"`
}
