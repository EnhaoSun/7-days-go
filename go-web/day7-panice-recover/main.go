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
	r := goe.Default()
	r.GET("/", func(ctx *goe.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Goe</h1>")
	})

	r.GET("/panic", func(ctx *goe.Context) {
		names := []string{"goe"}
		ctx.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
