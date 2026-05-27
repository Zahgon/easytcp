package easytcp

func newRouter() *Router { _ = "STUB: not implemented"; return nil }

// Router is a router for incoming message.
// Router routes the message to its handler and middlewares.
type Router struct {
	// handlerMapper maps message's ID to handler.
	// Handler will be called around middlewares.
	handlerMapper map[interface{}]HandlerFunc

	// middlewaresMapper maps message's ID to a list of middlewares.
	// These middlewares will be called before the handler in handlerMapper.
	middlewaresMapper map[interface{}][]MiddlewareFunc

	// globalMiddlewares is a list of MiddlewareFunc.
	// globalMiddlewares will be called before the ones in middlewaresMapper.
	globalMiddlewares []MiddlewareFunc

	notFoundHandler HandlerFunc
}

// HandlerFunc is the function type for handlers.
type HandlerFunc func(ctx Context)

// MiddlewareFunc is the function type for middlewares.
// A common pattern is like:
//
//	var mf MiddlewareFunc = func(next HandlerFunc) HandlerFunc {
//		return func(ctx Context) {
//			next(ctx)
//		}
//	}
type MiddlewareFunc func(next HandlerFunc) HandlerFunc

var nilHandler HandlerFunc = func(ctx Context) {}

// handleRequest walks ctx through middlewares and handler,
// and returns response message.
func (r *Router) handleRequest(ctx Context) { _ = "STUB: not implemented"; return }

// append to global ones

// create the handlers stack

// and call the handlers stack

// wrapHandlers wraps handler and middlewares into a right order call stack.
// Makes something like:
//
//	var wrapped HandlerFunc = m1(m2(m3(handle)))
func (r *Router) wrapHandlers(handler HandlerFunc, middles []MiddlewareFunc) (wrapped HandlerFunc) {
	_ = "STUB: not implemented"
	return *new(HandlerFunc)
}

// register stores handler and middlewares for id.
func (r *Router) register(id interface{}, h HandlerFunc, m ...MiddlewareFunc) {
	_ = "STUB: not implemented"
	return
}

// registerMiddleware stores the global middlewares.
func (r *Router) registerMiddleware(m ...MiddlewareFunc) { _ = "STUB: not implemented"; return }

// printHandlers prints registered route handlers to console.
func (r *Router) printHandlers(addr string) { _ = "STUB: not implemented"; return }

// don't uppercase the header
// respect the "\n" of cell content

// sort ids

// add table row

// route handler

// global middleware

// route middleware

func (r *Router) setNotFoundHandler(handler HandlerFunc) { _ = "STUB: not implemented"; return }
