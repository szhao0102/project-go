package main

import (
	"fmt"
	"net/http"

	"szhao.com/gee"
)

func main() {
	server := gee.Create()
	server.Get("/", func (w http.ResponseWriter, req *http.Request) {
		fmt.Println(req.URL)
	})
	err := server.Run(":80")
	if err != nil {
		fmt.Print(err)
	}
}