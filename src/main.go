package main

import (
	"net/http"
	"chain"
	"handler"
	"fmt"
	_ "server"
)

func main() {
	c := chain.Chain()
	c.RegisterHandle("handler1", handler.H1())
	c.RegisterHandle("handler2", handler.H2())
	c.RegisterHandle("router", handler.Router())

	err := http.ListenAndServe("127.0.0.1:8080", c)
	if err != nil {
		fmt.Println(err)
		fmt.Println("hello")
	}
}
