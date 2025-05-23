package events

type OrderCreatedEvent struct {
	OrderID     string  `json:"order_id"`
	CustomerID  string  `json:"customer_id"`
	CouponCode  string  `json:"coupon_code"`
	TotalAmount float64 `json:"total_amount"`
}
