package main

/*
curl "http://localhost:9999/hello/goektutu"
hello goektutu, you're at /hello/goektutu

curl "http://localhost:9999/assets/css/goektutu.css"
{"filepath":"css/goektutu.css"}
*/

import (
	"goe"
	"log"
	"net/http"
	"time"
)

func onlyForV2() goe.HandleFunc {
	return func(ctx *goe.Context) {
		// Start timer
		t := time.Now()
		// Process request
		ctx.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", ctx.StatusCode, ctx.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := goe.New()
	r.Use(goe.Logger()) // global middleware
	r.GET("/", func(ctx *goe.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Goe</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(c *goe.Context) {
			// expect /hello/goe
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
