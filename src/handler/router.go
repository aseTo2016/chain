package handler

import (
	"fmt"
	"chain"
	"net/http"
)

var router *Route

func init() {
	router = new(Route)
	router.routeInstances = make(map[string]map[string]RouterFunc, 10)
}

func Router() *Route{
	return router
}

type RouterFunc func(ctx *chain.Context)


type Route struct {
	routeInstances map[string]map[string]RouterFunc
}

func (r *Route)Handle(context * chain.Context) {
	r.router(context)
}

func (r *Route)AddRouter(method string, path string, rf RouterFunc) {
	if _, ok := r.routeInstances[method]; !ok {
		r.routeInstances[method] = make(map[string]RouterFunc, 10)
	}
	if _, ok := r.routeInstances[method][path]; ok {
		fmt.Println(fmt.Sprintf("method %s, path %s more registerd", method, path))
		return
	}
	r.routeInstances[method][path] = rf
	fmt.Println(fmt.Sprintf("method %s, path %s more register successfully", method, path))
}

func (r *Route)router(context *chain.Context) {
	req := context.Request()
	method := req.Method
	path := req.URL.Path
	fmt.Println(r.routeInstances[method])
	if _, ok := r.routeInstances[method]; !ok {
		context.WriteHeader(http.StatusNotFound)
		return
	}

	f, ok := r.routeInstances[method][path]
	if !ok {
		context.WriteHeader(http.StatusNotFound)
		return
	}
	f(context)
}