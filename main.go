package main

import (
	"fmt"
	"github.com/jiamengwei/gdd-space/gdd"
)

type Dog struct {
	Name string
	Age  int
}

func main() {
	g := gdd.New()

	g.GET("/", func(c *gdd.Context) {
		name := c.QueryValue("username")
		hello := fmt.Sprintf("hello %s", name)
		c.TEXT(hello)
	})

	g.POST("/addWithFormData", func(c *gdd.Context) {
		//获取参数
		username := c.PostFormValue("username")
		data := map[string]string{
			"username": username,
			"sex":      "man",
			"age":      "1",
		}
		c.JSON(data)
	})

	g.POST("/upload", func(c *gdd.Context) {
		c.FormFile("file")
	})

	g.POST("/addWithJSON", func(c *gdd.Context) {
		//获取参数
		var dog Dog
		c.Body(&dog)
		fmt.Printf("json body : %+v", dog)
		c.JSON(dog)
	})

	g.Run(":9999")
}
