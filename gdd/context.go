package gdd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

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

/**
获取参数
*/

func (c *Context) QueryValue(param string) string {
	return c.Request.URL.Query().Get(param)
}

func (c *Context) PostFormValue(param string) string {
	return c.Request.PostFormValue(param)
}

func (c *Context) FormValue(param string) string {
	return c.Request.FormValue(param)
}

func (c *Context) Body(data interface{}) error {
	//获取参数
	jsonDecoder := json.NewDecoder(c.Request.Body)
	err := jsonDecoder.Decode(data)
	if err != nil {
		fmt.Println(err)
		c.JSON(ERR("请求失败").SetStatus(http.StatusBadRequest))
		return err
	}
	return nil
}

func (c *Context) FormFile(fileParam string) {
	file, header, err := c.Request.FormFile(fileParam)
	if err != nil {
		fmt.Println(err)
		c.Body("file error")
		return
	}

	all, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(all)
		c.Body("file error")
		return
	}
	fmt.Println("file name: ", header.Filename)
	err = ioutil.WriteFile("filename.png", all, 0644)
	if err != nil {
		fmt.Println(err)
		c.Body("file error")
		return
	}
}

/**
封装响应
*/

func (c *Context) JSON(data interface{}) {
	c.RespWriter.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(c.RespWriter)
	apiResponse := OK(data)
	encoder.Encode(apiResponse)
}

func (c *Context) TEXT(text string) {
	c.RespWriter.Write([]byte(text))
}
