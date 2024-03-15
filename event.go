package eventbus

import (
	"errors"
)

type Event string

func (e Event) valid() error {
	if e == "" {
		return errors.New("event name can not be empty")
	}

	return nil
}
