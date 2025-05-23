
package events

import (
    "context"
    "encoding/json"
    "github.com/segmentio/kafka-go"
    "github.com/your-org/shared/events"
)

func publishOrderCreated(evt events.OrderCreatedEvent) {
    w := kafka.Writer{
        Addr:     kafka.TCP("kafka:9092"),
        Topic:    "order.created",
        Balancer: &kafka.LeastBytes{},
    }
    defer w.Close()
    data, _ := json.Marshal(evt)
    w.WriteMessages(context.Background(), kafka.Message{Value: data})
}
