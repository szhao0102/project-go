package gee

import (
	"net/http"
)

type HandlerFunc func(*Request, *Response)

type Engine struct {
	router *router
}

func Create() *Engine {
	return &Engine{router: initRouter()}
}

func (e *Engine) Get(uri string, handler HandlerFunc) {
	e.router.addRoute("GET", uri, handler)
}

func (e *Engine) Post(uri string, handler HandlerFunc) {
	e.router.addRoute("POST", uri, handler)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	newReq := &Request{httpReq: req, Method: req.Method, Path: req.URL.Path}
	newRes := &Response{w: w}
	e.router.handle(newReq, newRes)
}