package chain

import "net/http"

type Context struct {
	req *http.Request
	response http.ResponseWriter
	elements map[string]interface{}
	isStop bool
	status int
}

func NewContext(resp http.ResponseWriter, r *http.Request) *Context {
	c := new(Context)
	c.response =resp
	c.req = r
	c.elements = make(map[string]interface{}, 100)
	return c
}

func (c *Context) StopChain() {
	c.isStop = true
}

func (c *Context) IsStop() bool {
	return c.isStop
}

func (c *Context) Header() http.Header{
	return c.response.Header()
}

func (c *Context)WriteHeader(status int) {
	c.status = status
	if c.status < http.StatusOK || c.status > http.StatusMultipleChoices{
		c.StopChain()
	}
	c.response.WriteHeader(status)
}

func (c *Context) Write(data []byte) (int, error){
	return c.response.Write(data)
}

func (c *Context)Elements() map[string]interface{} {
	return c.elements
}

func (c *Context) Request() *http.Request {
	return c.req
}

//不支持并发
func (c *Context) AddElement(key string, value interface{}) {
	if len(key) == 0 {
		return
	}
	if c.elements == nil {
		c.elements = make(map[string]interface{}, 100)
	}
	c.elements[key] = value
}

type Handler interface {
	Handle(context * Context)
}

