package handlers

import (
	"fmt"

	"github.com/jaccovdalfsen/bestelcoupon/shared/events"
)

func handleOrderCreated(evt events.OrderCreatedEvent) {
	fmt.Printf("Coupon %s gebruikt voor order %s\n", evt.CouponCode, evt.OrderID)
}
