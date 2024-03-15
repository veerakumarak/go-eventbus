package main

import (
	"encoding/json"
	"fmt"
	"github.com/veerakumarak/go-eventbus"
	"time"
)

// Define commands
const (
	CreatedEvent eventbus.Event = "created_event"
	DeletedEvent                = "deleted_event"
)

// Define command payload
type CreatedPayload struct {
	Message string
}

type DeletedPayload struct {
	Name string
}

// Define subscribers
func Subscriber1(payload json.RawMessage) error {
	// parse payload
	var createdPayload CreatedPayload
	if err := json.Unmarshal(payload, &createdPayload); err != nil {
		return err
	}

	// execute the command
	time.Sleep(time.Second * 2)
	fmt.Println(createdPayload)

	// return the result
	return nil
}

func Subscriber2(payload json.RawMessage) error {
	// parse payload
	var deletedPayload DeletedPayload
	if err := json.Unmarshal(payload, &deletedPayload); err != nil {
		return err
	}

	// execute the command
	fmt.Println(deletedPayload)

	// return the result
	return nil
}

func main() {
	bus := eventbus.NewWithOptions("event-bus-2-100", 1, 100)

	// Register command handlers
	err := bus.Subscribe(CreatedEvent, Subscriber1)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = bus.Subscribe(DeletedEvent, Subscriber2)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Publish events
	payload, _ := json.Marshal(&CreatedPayload{Message: "user 1 created"})
	_ = bus.Publish(CreatedEvent, payload)

	payload, _ = json.Marshal(&DeletedPayload{Name: "user 1"})
	_ = bus.Publish(DeletedEvent, payload)

	payload, _ = json.Marshal(&CreatedPayload{Message: "user 2 created"})
	_ = bus.Publish(CreatedEvent, payload)

	payload, _ = json.Marshal(&DeletedPayload{Name: "user 2"})
	_ = bus.Publish(DeletedEvent, payload)

	fmt.Println("published events")
	bus.Shutdown()
}
