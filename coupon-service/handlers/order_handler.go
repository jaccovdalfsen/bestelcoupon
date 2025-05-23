
package handlers

import "fmt"
import "github.com/your-org/shared/events"

func handleOrderCreated(evt events.OrderCreatedEvent) {
    fmt.Printf("Coupon %s gebruikt voor order %s\n", evt.CouponCode, evt.OrderID)
}
