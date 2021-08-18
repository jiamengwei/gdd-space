package main

import (
	"github.com/jiamengwei/gdd-space/gdd"
)

func main() {
	g := gdd.New()

	g.GET("/", func(c *gdd.Context) {
		name := c.Query("name")
		c.RespWriter.Write([]byte("hello" + name))
	})

	g.Run(":9999")

}
