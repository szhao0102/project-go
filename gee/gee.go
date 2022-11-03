package gee

import "net/http"

type HandlerFunc func(http.ResponseWriter, *http.Request) 

type Engine struct {
	router map[string]HandlerFunc
}

func Create() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (e *Engine) Get(uri string, handler HandlerFunc) {
	e.router["get_" + uri] = handler
}

func (e *Engine) Post(uri string, handler HandlerFunc) {
	e.router["post_"+uri] = handler
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "_" + req.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, req)
	}
}