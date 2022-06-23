package main

/*
curl "http://localhost:9999/hello/goektutu"
hello goektutu, you're at /hello/goektutu

curl "http://localhost:9999/assets/css/goektutu.css"
{"filepath":"css/goektutu.css"}
*/

import (
	"goe"
	"net/http"
)

func main() {
	r := goe.New()
	r.GET("/", func(ctx *goe.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Goe</h1>")
	})

	r.GET("/hello", func(ctx *goe.Context) {
		// expect /hello?name=goe
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
	})

	r.GET("/hello/:name", func(c *goe.Context) {
		// expect /hello/goe
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *goe.Context) {
		c.JSON(http.StatusOK, goe.O{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
