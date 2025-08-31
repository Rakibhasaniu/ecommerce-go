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

// func (mngr *Manager) Use(middlewares ...Middleware) *Manager {
// 	mngr.globalMiddleware = append(mngr.globalMiddleware, middlewares...)
// 	return mngr
// }

func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	n := next
	// for i := len(middlewares) - 1; i >= 0; i-- {
	// 	middleware := middlewares[i]
	// 	n = middleware(n)
	// }

	for _, middleware := range middlewares {
		n = middleware(n)
	}
	return n

}
