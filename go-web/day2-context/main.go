package main

/*
(1)
$ curl -i http://localhost:9999/
HTTP/1.1 200 OK
Date: Mon, 12 Aug 2019 16:52:52 GMT
Content-Length: 18
Content-Type: text/html; charset=utf-8
<h1>Hello Gee</h1>

(2)
$ curl "http://localhost:9999/hello?name=geektutu"
hello geektutu, you're at /hello

(3)
$ curl "http://localhost:9999/login" -X POST -d 'username=geektutu&password=1234'
{"password":"1234","username":"geektutu"}

(4)
$ curl "http://localhost:9999/xxx"
404 NOT FOUND: /xxx
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

	r.POST("/login", func(ctx *goe.Context) {
		ctx.JSON(http.StatusOK, goe.O{
			"username": ctx.PostForm("username"),
			"password": ctx.PostForm("password"),
		})
	})

	r.Run(":9999")
}
