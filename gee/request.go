package gee

import (
	"io"
	"net/http"
)

type Request struct {
	httpReq *http.Request
}

func (req *Request) Method() string {
	return req.httpReq.Method
}

func (req *Request) Params(key string) string {
	return req.httpReq.URL.Query().Get(key)
}

func (req *Request) Body() (io.ReadCloser, error){
	return req.httpReq.GetBody()
}

