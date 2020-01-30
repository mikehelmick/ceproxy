package main

import (
	"context"
	"log"
	"os"

	cloudevents "github.com/cloudevents/sdk-go"
	"knative.dev/eventing/pkg/kncloudevents"
)

func main() {
	brokerAddr := os.Getenv("broker")
	if brokerAddr == "" {
		brokerAddr = "http://default-broker.default.svc.cluster.local"
		log.Printf("Using default broker addr: %s", brokerAddr)
	}

	c, err := kncloudevents.NewDefaultClient()
	if err != nil {
		log.Fatal("Failed to create client, ", err)
	}

	log.Fatal(c.StartReceiver(context.Background(),
		func(event cloudevents.Event) {
			log.Printf("Proxying event: %v", event)
			ctx := cloudevents.ContextWithTarget(context.Background(), brokerAddr)
			_, _, err := c.Send(ctx, event)
			if err != nil {
				log.Fatalf("Error proxying event: %v | error: %v", event, err)
			}
		}))
}
