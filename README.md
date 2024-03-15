# Event Bus Library in Go

This library provides a lightweight and flexible implementation of an Event Bus in Go. It allows you to publish and subscribe to events in a decoupled manner, enabling loose coupling between components in your application.

## Features

- **Decoupling**: Allows components to communicate with each other without direct dependencies, promoting better code organization and testability.
- **Flexibility**: Provides a simple and extensible API for publishing and subscribing to events.
- **Concurrency**: events asynchronously, making it suitable for high-performance applications.

## Installation

```bash
go get github.com/veerakumarak/go-eventbus
```

## Usage

<pre><div class="dark bg-gray-950 rounded-md"><div class="flex items-center relative text-token-text-secondary bg-token-main-surface-secondary px-4 py-2 text-xs font-sans justify-between rounded-t-md"><span></span><span class="" data-state="closed"><button class="flex gap-1 items-center"><svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="icon-sm"><path fill-rule="evenodd" clip-rule="evenodd" d="M12 4C10.8954 4 10 4.89543 10 6H14C14 4.89543 13.1046 4 12 4ZM8.53513 4C9.22675 2.8044 10.5194 2 12 2C13.4806 2 14.7733 2.8044 15.4649 4H17C18.6569 4 20 5.34315 20 7V19C20 20.6569 18.6569 22 17 22H7C5.34315 22 4 20.6569 4 19V7C4 5.34315 5.34315 4 7 4H8.53513ZM8 6H7C6.44772 6 6 6.44772 6 7V19C6 19.5523 6.44772 20 7 20H17C17.5523 20 18 19.5523 18 19V7C18 6.44772 17.5523 6 17 6H16C16 7.10457 15.1046 8 14 8H10C8.89543 8 8 7.10457 8 6Z" fill="currentColor"></path></svg></button></span></div><div class="p-4 overflow-y-auto"><code class="!whitespace-pre hljs language-go">
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
</code></div></div></pre>

### The output will be

```
{user 1 created}
{user 1}
{user 2 created}
{user 2}
```

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request for any improvements or new features you'd like to see.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
