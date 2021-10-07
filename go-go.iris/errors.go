package main

import "github.com/kataras/iris/v12"

func NotFound(ctx iris.Context) {
	// 当出现 404 时， 渲染模版
	// $views_dir/errors/404.html
	ctx.View("errors/404.html")
}

func InternalServerError(ctx iris.Context) {
	ctx.WriteString("Oups something went wrong, try again")
}
