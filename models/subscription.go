package models

type Subscription struct {
	PlanID     string  `json:"plan_id"`
	CustomerID string  `json:"customer_id"`
	Price      float64 `json:"price"`
}
