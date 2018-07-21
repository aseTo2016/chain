package chain

import (
	"net/http"
	"sync"
	"fmt"
)

var c *chain

func init() {
	c = &chain{}
	c.handlerNames = make(map[string]interface{}, 10)
	c.handlers = make([]Handler, 0, 10)
}

type chain struct {
	handlerNames map[string]interface{}
	handlers []Handler
	lock sync.Mutex
}

func Chain() *chain {
	return c
}



func (c *chain)RegisterHandle(name string, h Handler) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if _, ok := c.handlerNames[name]; ok {
		fmt.Println(fmt.Sprintf("handler %s more register", name))
		return
	}
	c.handlerNames[name] = name
	c.handlers = append(c.handlers, h)
	fmt.Println(fmt.Sprintf("register handler %s successfully", name))
}

func (c *chain)ServeHTTP(resp http.ResponseWriter, r *http.Request) {
	ctx := NewContext(resp, r)
	for _, h := range c.handlers {
		h.Handle(ctx)
		if ctx.IsStop() {
			return
		}
	}
}