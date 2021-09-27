package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	// Creates a iris app with default middleware:
	// logger and recovery (crash-free) middleware
	app := iris.Default()
	// custom middleware
	app.Use(myMiddleware)
	app.Get("/", index)

	app.Get("/assets/{asset:path}", func(ctx iris.Context) {
		ctx.JSON(ctx.Params().Get("asset"))
	})

	app.Get("/user/{name:string}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.JSON(iris.Map{
			"name": name,
		})
	})

	app.Get("/user/{userid:int min(1)}", func(ctx iris.Context) {
		id := ctx.Params().Get("userid")
		ctx.JSON(iris.Map{
			"id": id,
		})
	})

	app.Get("{root:path}", func(ctx iris.Context) {
		root := ctx.Params().Get("root")
		ctx.JSON(iris.Map{
			"root": root,
		})
	})

	app.Listen(":8080")
}

func index(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"message": "hello",
	})
}

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
