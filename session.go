package easytcp

import (
	"net"
	"sync"
	"time"
)

// Session represents a TCP session.
type Session interface {
	// ID returns current session's id.
	ID() interface{}

	// SetID sets current session's id.
	SetID(id interface{})

	// Send sends the ctx to the respStream.
	Send(ctx Context) bool

	// Codec returns the codec, can be nil.
	Codec() Codec

	// Close closes current session.
	Close()

	// AllocateContext gets a Context ships with current session.
	AllocateContext() Context

	// Conn returns the underlined connection.
	Conn() net.Conn

	// AfterCreateHook blocks until session's on-create hook triggered.
	AfterCreateHook() <-chan struct{}

	// AfterCloseHook blocks until session's on-close hook triggered.
	AfterCloseHook() <-chan struct{}
}

type session struct {
	id               interface{}   // session's ID.
	conn             net.Conn      // tcp connection
	closedC          chan struct{} // to close when read/write loop stopped
	closeOnce        sync.Once     // ensure one session only close once
	afterCreateHookC chan struct{} // to close after session's on-create hook triggered
	afterCloseHookC  chan struct{} // to close after session's on-close hook triggered
	respStream       chan Context  // response queue channel, pushed in Send() and popped in writeOutbound()
	packer           Packer        // to pack and unpack message
	codec            Codec         // encode/decode message data
	ctxPool          sync.Pool     // router context pool
	asyncRouter      bool          // calls router HandlerFunc in a goroutine if false
}

// sessionOption is the extra options for session.
type sessionOption struct {
	Packer        Packer
	Codec         Codec
	respQueueSize int
	asyncRouter   bool
}

// newSession creates a new session.
// Parameter conn is the TCP connection,
// opt includes packer, codec, and channel size.
// Returns a session pointer.
func newSession(conn net.Conn, opt *sessionOption) *session { _ = "STUB: not implemented"; return nil }

// use uuid as default

// ID returns the session's id.
func (s *session) ID() interface{} {
	_ = "STUB: not implemented"

	// SetID sets session id.
	// Can be called in server.OnSessionCreate() callback.
	return nil
}

func (s *session) SetID(id interface{}) {
	_ = "STUB: not implemented"

	// Send pushes response message to respStream.
	// Returns false if session is closed or ctx is done.
	return
}

func (s *session) Send(ctx Context) (ok bool) { _ = "STUB: not implemented"; return false }

// Codec implements Session Codec.
func (s *session) Codec() Codec {
	_ = "STUB: not implemented"

	// Close closes the session, but doesn't close the connection.
	// The connection will be closed in the server once the session's closed.
	return *new(Codec)
}

func (s *session) Close() { _ = "STUB: not implemented"; return }

// AfterCreateHook blocks until session's on-create hook triggered.
func (s *session) AfterCreateHook() <-chan struct{} { _ = "STUB: not implemented"; return nil }

// AfterCloseHook blocks until session's on-close hook triggered.
func (s *session) AfterCloseHook() <-chan struct{} { _ = "STUB: not implemented"; return nil }

// AllocateContext gets a Context from pool and reset all but session.
func (s *session) AllocateContext() Context { _ = "STUB: not implemented"; return *new(Context) }

// Conn returns the underlined connection instance.
func (s *session) Conn() net.Conn {
	_ = "STUB: not implemented"

	// readInbound reads message packet from connection in a loop.
	// And send unpacked message to reqQueue, which will be consumed in router.
	// The loop breaks if errors occurred or the session is closed.
	return *new(net.Conn)
}

func (s *session) readInbound(router *Router, timeout time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (s *session) handleReq(router *Router, reqMsg *Message) { _ = "STUB: not implemented"; return }

// writeOutbound fetches message from respStream channel and writes to TCP connection in a loop.
// Parameter writeTimeout specified the connection writing timeout.
// The loop breaks if errors occurred, or the session is closed.
func (s *session) writeOutbound(writeTimeout time.Duration) { _ = "STUB: not implemented"; return }

func (s *session) packResponse(ctx Context) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
