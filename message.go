package easytcp

import (
	"sync"
)

// NewMessage creates a Message pointer.
func NewMessage(id interface{}, data []byte) *Message { _ = "STUB: not implemented"; return nil }

// Message is the abstract of inbound and outbound message.
type Message struct {
	id      interface{}
	data    []byte
	storage map[string]interface{}
	mu      sync.RWMutex
}

// ID returns the id of current message.
func (m *Message) ID() interface{} {
	_ = "STUB: not implemented"

	// Data returns the data part of current message.
	return nil
}

func (m *Message) Data() []byte {
	_ = "STUB: not implemented"

	// Set stores kv pair.
	return nil
}

func (m *Message) Set(key string, value interface{}) { _ = "STUB: not implemented"; return }

// Get retrieves the value according to the key.
func (m *Message) Get(key string) (value interface{}, exists bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// MustGet retrieves the value according to the key.
// Panics if key does not exist.
func (m *Message) MustGet(key string) interface{} { _ = "STUB: not implemented"; return nil }

// Remove deletes the key from storage.
func (m *Message) Remove(key string) { _ = "STUB: not implemented"; return }
