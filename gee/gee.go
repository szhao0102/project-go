package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(*Request, *Response)

type Engine struct {
	router map[string]HandlerFunc
}

func Create() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (e *Engine) Get(uri string, handler HandlerFunc) {
	e.router["GET_"+uri] = handler
}

func (e *Engine) Post(uri string, handler HandlerFunc) {
	e.router["POST_"+uri] = handler
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "_" + req.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(&Request{httpReq: req}, &Response{w: w})
	} else {
		fmt.Fprintf(w, "404 %s\n", req.URL)
	}
}