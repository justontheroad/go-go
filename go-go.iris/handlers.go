package main

import "github.com/kataras/iris/v12"

func Before(ctx iris.Context) {
	info := "Welcome"
	requestPath := ctx.Path()
	println("Before the mainHandler: " + requestPath)

	ctx.Values().Set("info", info)
	ctx.Next()
}

func After(ctx iris.Context) {
	println("After the mainHandler")
}

func MainHandler(ctx iris.Context) {
	println("Inside mainHandler")

	info := ctx.Values().GetString("info")

	// 向客户端写一些内容作为响应。
	ctx.HTML("<h1>Response</h1>")
	ctx.HTML("<br/> Info: " + info)

	ctx.Next() // 执行 "after" 中间件
}

func IndexHandler(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"message": "hello",
	})
}

func CustomHandler(ctx iris.Context) {
	ctx.Writef("Hello from method: %s and path: %s\n", ctx.Method(), ctx.Path())
}

func UserProfileHandler(ctx iris.Context) {
	ctx.Writef("user %s profile\n", ctx.Params().Get("id"))
}

func UserMessageHandler(ctx iris.Context) {
	ctx.Writef("user %s messge\n", ctx.Params().Get("id"))
}

func OtherHandler(ctx iris.Context) {
	ctx.Writef("custom router wrapper")
}
