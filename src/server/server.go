package server

import (
	"chain"
	"fmt"
	"handler"
	"net/http"
)

func init() {
	handler.Router().AddRouter(http.MethodGet, "/hello", Hello)
	handler.Router().AddRouter(http.MethodGet, "/aseto", Aseto)
}

func Hello(ctx *chain.Context) {
	fmt.Println(ctx.Elements())
	fmt.Println("hello")
	ctx.Write([]byte("hello"))
}


func Aseto(ctx *chain.Context) {
	fmt.Println(ctx.Elements())
	fmt.Println("hello")
	ctx.Write([]byte("hello"))
}