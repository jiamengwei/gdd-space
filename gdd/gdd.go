package gdd

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	routerGroups RouterGroup
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
		routerGroups: RouterGroup{
			Group:  map[string]string{},
			Router: map[string]HandlerFunc{},
		},
	}
}

func (e *Engine) Group(prefix string) *RouterGroup {
	e.routerGroups.NewGroup(prefix)
	e.routerGroups.Engine = e
	return &e.routerGroups

}

func (e *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	method := request.Method
	key := fmt.Sprintf("%s:%s", method, path)

	if handler, ok := e.routerGroups.Router[key]; ok {
		context := new(writer, request)
		handler(context)
		return
	}
	writer.WriteHeader(http.StatusNotFound)
	writer.Write([]byte("404 not found"))
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute(http.MethodGet, pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRoute(http.MethodPost, pattern, handler)
}

func (e *Engine) addRoute(method, pattern string, handler HandlerFunc) {
	key := fmt.Sprintf("%s:%s", method, pattern)
	e.routerGroups.Router[key] = handler
}
