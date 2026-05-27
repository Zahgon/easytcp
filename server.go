package easytcp

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

//go:generate mockgen -destination internal/mock/server_mock.go -package mock net Listener,Error,Conn

// Server is a server for TCP connections.
type Server struct {
	Listener net.Listener

	// Packer is the message packer, will be passed to session.
	Packer Packer

	// Codec is the message codec, will be passed to session.
	Codec Codec

	// OnSessionCreate is an event hook, will be invoked when session's created.
	OnSessionCreate func(sess Session)

	// OnSessionClose is an event hook, will be invoked when session's closed.
	OnSessionClose func(sess Session)

	socketReadBufferSize  int
	socketWriteBufferSize int
	socketSendDelay       bool
	readTimeout           time.Duration
	writeTimeout          time.Duration
	respQueueSize         int
	router                *Router
	printRoutes           bool
	acceptingC            chan struct{}
	stoppedC              chan struct{}
	asyncRouter           bool
}

// ServerOption is the option for Server.
type ServerOption struct {
	SocketReadBufferSize  int           // sets the socket read buffer size.
	SocketWriteBufferSize int           // sets the socket write buffer size.
	SocketSendDelay       bool          // sets the socket delay or not.
	ReadTimeout           time.Duration // sets the timeout for connection read.
	WriteTimeout          time.Duration // sets the timeout for connection write.
	Packer                Packer        // packs and unpacks packet payload, default packer is the DefaultPacker.
	Codec                 Codec         // encodes and decodes the message data, can be nil.
	RespQueueSize         int           // sets the response channel size of session, DefaultRespQueueSize will be used if < 0.
	DoNotPrintRoutes      bool          // whether to print registered route handlers to the console.

	// AsyncRouter represents whether to execute a route HandlerFunc of each session in a goroutine.
	// true means execute in a goroutine.
	AsyncRouter bool
}

// ErrServerStopped is returned when server stopped.
var ErrServerStopped = fmt.Errorf("server stopped")

const DefaultRespQueueSize = 1024

// NewServer creates a Server according to opt.
func NewServer(opt *ServerOption) *Server { _ = "STUB: not implemented"; return nil }

// Serve starts to serve the lis.
func (s *Server) Serve(lis net.Listener) error { _ = "STUB: not implemented"; return nil }

// Run starts to listen TCP and keeps accepting TCP connection in a loop.
// The loop breaks when error occurred, and the error will be returned.
func (s *Server) Run(addr string) error { _ = "STUB: not implemented"; return nil }

// RunTLS starts serve TCP with TLS.
func (s *Server) RunTLS(addr string, config *tls.Config) error {
	_ = "STUB: not implemented"
	return nil
}

// acceptLoop accepts TCP connections in a loop, and handle connections in goroutines.
// Returns error when error occurred.
func (s *Server) acceptLoop() error { _ = "STUB: not implemented"; return nil }

// handleConn creates a new session with `conn`,
// handles the message through the session in different goroutines,
// and waits until the session's closed, then close the `conn`.
func (s *Server) handleConn(conn net.Conn) {
	_ = "STUB: not implemented"
	// nolint
	return
}

// start reading message packet from connection.
// start writing message packet to connection.

// wait for session finished.
// or the server is stopped.

// Stop stops server. Closing Listener and all connections.
func (s *Server) Stop() error { _ = "STUB: not implemented"; return nil }

// AddRoute registers message handler and middlewares to the router.
func (s *Server) AddRoute(msgID interface{}, handler HandlerFunc, middlewares ...MiddlewareFunc) {
	_ = "STUB: not implemented"
	return
}

// Use registers global middlewares to the router.
func (s *Server) Use(middlewares ...MiddlewareFunc) { _ = "STUB: not implemented"; return }

// NotFoundHandler sets the not-found handler for router.
func (s *Server) NotFoundHandler(handler HandlerFunc) { _ = "STUB: not implemented"; return }

func (s *Server) isStopped() bool { _ = "STUB: not implemented"; return false }
