package main

import (
	"fmt"
	"szhao.com/gee"
)

func main() {
	server := gee.Create()
	server.Get("/", func(req *gee.Request, res *gee.Response) {
		fmt.Println(req.Method)
	})
	err := server.Run(":80")
	if err != nil {
		fmt.Print(err)
	}
}