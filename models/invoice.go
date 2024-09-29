package models

type Invoice struct {
	InvoiceID  string  `json:"invoice_id"`
	MerchantID string  `json:"merchant_id"`
	Amount     float64 `json:"amount"`
	Status     string  `json:"status"` // e.g., "Pending", "Paid", "Cancelled"
}
