
package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/your-org/shared/events"
)

func main() {
    http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
        var evt events.OrderCreatedEvent
        json.NewDecoder(r.Body).Decode(&evt)
        publishOrderCreated(evt)
        w.WriteHeader(http.StatusAccepted)
    })
    log.Fatal(http.ListenAndServe(":8080", nil))
}
