package gee

import "fmt"

type router struct {
	handlers map[string]HandlerFunc
}

func initRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method, pattern string, fun HandlerFunc) {
	r.handlers[method + "_" + pattern] = fun
}

func (r *router) handle(req *Request, res *Response) {
	key := req.Method + "_" + req.Path
	if handler, ok := r.handlers[key]; ok {
		handler(req, res)
	} else {
		fmt.Println("404", req.Path)
	}
}