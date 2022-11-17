package gee

import (
	"io"
	"net/http"
)

type Request struct {
	httpReq *http.Request
	Method  string
	Path string
}

func (req *Request) GetMethod() string {
	return req.Method
}

func (req *Request) GetPath() string {
	return req.Path
}

func (req *Request) Params(key string) string {
	return req.httpReq.URL.Query().Get(key)
}

func (req *Request) Body() (io.ReadCloser, error){
	return req.httpReq.GetBody()
}

