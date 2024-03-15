package eventbus

import "encoding/json"

type HandlerFunc func(payload json.RawMessage) error

func (h HandlerFunc) valid() error {
	return nil
}
