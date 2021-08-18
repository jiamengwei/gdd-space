package gdd

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	routers map[string]HandlerFunc
}

func (e *Engine) Run(addr string) {
	fmt.Println("web server running at ", addr)
	err := http.ListenAndServe(addr, e)
	if err != nil {
		panic(err)
	}
}

func New() *Engine {
	return &Engine{
		routers: make(map[string]HandlerFunc),
	}
}

func (e *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	method := request.Method
	key := fmt.Sprintf("%s:%s", method, path)

	if handler, ok := e.routers[key]; ok {
		context := new(writer, request)
		handler(context)
		return
	}
	writer.WriteHeader(http.StatusNotFound)
	writer.Write([]byte("404 not found"))
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

func (e *Engine) addRoute(method, pattern string, handler HandlerFunc) {
	key := fmt.Sprintf("%s:%s", method, pattern)
	e.routers[key] = handler
}
