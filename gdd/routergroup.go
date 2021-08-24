package gdd

import "net/http"

type RouterGroup struct {
	Prefix string
	Group  map[string]string
	Router map[string]HandlerFunc
	*Engine
}

func (rg *RouterGroup) NewGroup(prefix string) {
	rg.Prefix = prefix
}

func (rg *RouterGroup) GET(pattern string, handler HandlerFunc) {
	pattern = rg.Prefix + pattern
	rg.Group[pattern] = rg.Prefix
	rg.addRoute(http.MethodGet, pattern, handler)
}

func (rg *RouterGroup) POST(pattern string, handler HandlerFunc) {
	pattern = rg.Prefix + pattern
	rg.addRoute(http.MethodPost, pattern, handler)
}
