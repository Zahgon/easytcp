package easytcp

import (
	"context"
	"sync"
	"time"
)

// Context is a generic context in a message routing.
// It allows us to pass variables between handler and middlewares.
type Context interface {
	context.Context

	// WithContext sets the underline context.
	// It's very useful to control the workflow when send to response channel.
	WithContext(ctx context.Context) Context

	// Session returns the current session.
	Session() Session

	// SetSession sets session.
	SetSession(sess Session) Context

	// Request returns request message.
	Request() *Message

	// SetRequest encodes data with session's codec and sets request message.
	SetRequest(id, data interface{}) error

	// MustSetRequest encodes data with session's codec and sets request message.
	// panics on error.
	MustSetRequest(id, data interface{}) Context

	// SetRequestMessage sets request message directly.
	SetRequestMessage(msg *Message) Context

	// Bind decodes request message to v.
	Bind(v interface{}) error

	// Response returns the response message.
	Response() *Message

	// SetResponse encodes data with session's codec and sets response message.
	SetResponse(id, data interface{}) error

	// MustSetResponse encodes data with session's codec and sets response message.
	// panics on error.
	MustSetResponse(id, data interface{}) Context

	// SetResponseMessage sets response message directly.
	SetResponseMessage(msg *Message) Context

	// Send sends itself to current session.
	Send() bool

	// SendTo sends itself to session.
	SendTo(session Session) bool

	// Get returns key value from storage.
	Get(key string) (value interface{}, exists bool)

	// Set store key value into storage.
	Set(key string, value interface{})

	// Remove deletes the key from storage.
	Remove(key string)

	// Copy returns a copy of Context.
	Copy() Context
}

var _ Context = &routeContext{} // implementation check

// newContext creates a routeContext pointer.
func newContext() *routeContext { _ = "STUB: not implemented"; return nil }

// routeContext implements the Context interface.
type routeContext struct {
	rawCtx  context.Context
	mu      sync.RWMutex
	storage map[string]interface{}
	session Session
	reqMsg  *Message
	respMsg *Message
}

// Deadline implements the context.Context Deadline method.
func (c *routeContext) Deadline() (time.Time, bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

// Done implements the context.Context Done method.
func (c *routeContext) Done() <-chan struct{} { _ = "STUB: not implemented"; return nil }

// Err implements the context.Context Err method.
func (c *routeContext) Err() error { _ = "STUB: not implemented"; return nil }

// Value implements the context.Context Value method.
func (c *routeContext) Value(key interface{}) interface{} { _ = "STUB: not implemented"; return nil }

// WithContext sets the underline context.
func (c *routeContext) WithContext(ctx context.Context) Context {
	_ = "STUB: not implemented"
	return *

	// Session implements Context.Session method.
	new(Context)
}

func (c *routeContext) Session() Session {
	_ = "STUB: not implemented"

	// SetSession sets session.
	return *new(Session)
}

func (c *routeContext) SetSession(sess Session) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// Request implements Context.Request method.
func (c *routeContext) Request() *Message {
	_ = "STUB: not implemented"

	// SetRequest sets request by id and data.
	return nil
}

func (c *routeContext) SetRequest(id, data interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// MustSetRequest implements Context.MustSetRequest method.
func (c *routeContext) MustSetRequest(id, data interface{}) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// SetRequestMessage sets request message.
func (c *routeContext) SetRequestMessage(msg *Message) Context {
	_ = "STUB: not implemented"
	return *

	// Bind implements Context.Bind method.
	new(Context)
}

func (c *routeContext) Bind(v interface{}) error { _ = "STUB: not implemented"; return nil }

// Response implements Context.Response method.
func (c *routeContext) Response() *Message {
	_ = "STUB: not implemented"

	// SetResponse implements Context.SetResponse method.
	return nil
}

func (c *routeContext) SetResponse(id, data interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// MustSetResponse implements Context.MustSetResponse method.
func (c *routeContext) MustSetResponse(id, data interface{}) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// SetResponseMessage implements Context.SetResponseMessage method.
func (c *routeContext) SetResponseMessage(msg *Message) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

// Send implements Context.Send method.
func (c *routeContext) Send() bool { _ = "STUB: not implemented"; return false }

// SendTo implements Context.SendTo method.
func (c *routeContext) SendTo(sess Session) bool { _ = "STUB: not implemented"; return false }

// Get implements Context.Get method.
func (c *routeContext) Get(key string) (value interface{}, exists bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Set implements Context.Set method.
func (c *routeContext) Set(key string, value interface{}) { _ = "STUB: not implemented"; return }

// Remove implements Context.Remove method.
func (c *routeContext) Remove(key string) { _ = "STUB: not implemented"; return }

// Copy implements Context.Copy method.
func (c *routeContext) Copy() Context { _ = "STUB: not implemented"; return *new(Context) }

func (c *routeContext) reset() { _ = "STUB: not implemented"; return }
