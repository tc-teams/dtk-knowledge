package api

import (
	"fmt"
)

//Middleware define a function thats takes   in a http.Handler  and return a http.Handler
type MiddlewareFunc func(Handler) Handler
//MiddlewareContext for api
type Middleware struct {
	context []MiddlewareFunc
}

func(m *Middleware) Chain(mwa ...MiddlewareFunc){
	m.context = append(m.context, mwa...)
}

//Hello add a new middleware
func (m *Middleware) ChainMiddleware(h Handler) Handler {
	// if our chain is done, use the original handlerfunc
	if len(m.context) == 0 {
		return h
	}
	for _, adapter := range m.context {
		h = adapter(h)
	}
	fmt.Println("adicionando os middleware")
	return h
}

func newMiddlewareContext(mws ...MiddlewareFunc)  *Middleware {
	return &Middleware{
		context: mws,
	}
	
}

