package handler

import (
	"chain"
	"fmt"
)

func H2() chain.Handler {
	return &Handler2{}
}

type Handler2 struct {
}

func (h *Handler2)Handle(context * chain.Context) {
	fmt.Println(context.Elements())
}
