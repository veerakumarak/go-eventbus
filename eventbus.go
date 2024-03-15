package eventbus

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/veerakumarak/go-workerpool"
)

type IEventBus interface {
	Subscribe(Event, HandlerFunc) error
	Publish(Event, json.RawMessage) error
	Shutdown()
}

type bus struct {
	handlers map[Event][]HandlerFunc
	pool     workerpool.IWorkerPool
	quit     bool
}

func New(name string) IEventBus {
	pool := workerpool.New(name, 1, 1)
	pool.Start()
	return &bus{
		handlers: make(map[Event][]HandlerFunc),
		pool:     pool,
	}
}

func NewWithOptions(name string, maxWorkers int, queueSize int) IEventBus {
	pool := workerpool.New(name, maxWorkers, queueSize)
	pool.Start()
	return &bus{
		handlers: make(map[Event][]HandlerFunc),
		pool:     pool,
	}
}

func (b *bus) Subscribe(event Event, fn HandlerFunc) error {
	if err := event.valid(); err != nil {
		return err
	}

	if err := fn.valid(); err != nil {
		return err
	}

	b.handlers[event] = append(b.getHandlers(event), fn)

	return nil
}

func (b *bus) Publish(event Event, payload json.RawMessage) error {
	if b.quit {
		return errors.New("shutting down, cannot publish the event")
	}
	if err := b.validate(event, payload); err != nil {
		return err
	}
	return b.pool.Submit(createAsync(b, event, payload))
}

func (b *bus) Shutdown() {
	b.quit = true
	b.pool.Shutdown()
}

func createAsync(b *bus, event Event, payload json.RawMessage) workerpool.Task {
	return func() {
		_ = b.execute(event, payload)
	}
}

func (b *bus) validate(event Event, payload json.RawMessage) error {
	if err := event.valid(); err != nil {
		return err
	}

	return Message(payload).valid()
}

func (b *bus) execute(event Event, payload json.RawMessage) error {
	if err := b.validate(event, payload); err != nil {
		fmt.Println(err)
		return err
	}

	h := b.getHandlers(event)
	for _, handler := range h {
		handler(payload)
	}

	return nil
}

func (b *bus) getHandlers(event Event) []HandlerFunc {
	if h, ok := b.handlers[event]; ok {
		return h
	}
	b.handlers[event] = []HandlerFunc{}
	return b.handlers[event]
}
