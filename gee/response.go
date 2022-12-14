package gee

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	w http.ResponseWriter
}

func (res *Response) Send(content string) {
	res.w.WriteHeader(200)
	res.w.Write([]byte(content))
}

func (res *Response) Json(content any) {
	res.w.WriteHeader(200)
	res.w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(res.w)
	if err := encoder.Encode(content); err != nil {
		http.Error(res.w, err.Error(), 500)
	}
}
