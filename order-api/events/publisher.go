package events

import (
	"context"
	"encoding/json"

	"github.com/jaccovdalfsen/bestelcoupon/shared/events"
	"github.com/segmentio/kafka-go"
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
