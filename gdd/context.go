package gdd

import "net/http"

type Context struct {
	RespWriter http.ResponseWriter
	Request    *http.Request
}

func new(writer http.ResponseWriter, request *http.Request) *Context {
	return &Context{
		RespWriter: writer,
		Request:    request,
	}
}

func (c *Context) Query(param string) string {
	return c.Request.URL.Query().Get(param)
}

func (c *Context) JSON() {

}
