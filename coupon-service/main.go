
package main

import (
    "context"
    "encoding/json"
    "log"

    "github.com/segmentio/kafka-go"
    "github.com/your-org/shared/events"
)

func main() {
    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{"kafka:9092"},
        Topic:   "order.created",
        GroupID: "coupon-service",
    })
    for {
        m, _ := r.ReadMessage(context.Background())
        var evt events.OrderCreatedEvent
        json.Unmarshal(m.Value, &evt)
        handleOrderCreated(evt)
    }
}
