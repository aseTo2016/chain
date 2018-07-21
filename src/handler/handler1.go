package handler

import (
	"chain"
)

func H1() chain.Handler {
	return &Handler1{}
}

type Handler1 struct {
}

func (h *Handler1)Handle(context * chain.Context) {
	context.AddElement("h1", "h1")
}
