package models

type Product struct {
    MerchantID int     `json:"MerchantID"`
    ProductID  int     `json:"ProductID"`
    Name       string  `json:"Name"`
    Price      float64 `json:"Price"`
    Quantity   int     `json:"Quantity"`
}
