package model

type OrderCount struct {
	OrderId    string   `json:"orderId,omitempty"`
	Count      int      `json:"count,omitempty"`
}
