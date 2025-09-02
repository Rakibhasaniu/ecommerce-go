package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler
type Manager struct {
	globalMiddleware []Middleware
}

func NewManager() *Manager {
	mngr := Manager{
		globalMiddleware: make([]Middleware, 0),
	}
	return &mngr
}

func (mngr *Manager) Use(middlewares ...Middleware) *Manager {
	mngr.globalMiddleware = append(mngr.globalMiddleware, middlewares...)
	return mngr
}

func (mngr *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {
	h := handler
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}
func (mngr *Manager) wrappedMux(handler http.Handler, middlewares ...Middleware) http.Handler {
	h := handler
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}
