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
	// Use 和 Done 应用于当前路由分组和它的子分组，即在调用 Use 或 Done 之前的路由，不会应用该中间件
	app.Get("/home", indexHandler)
	app.Use(before)
	app.Done(after)
	// 在路由注册前使用 `app.Use/Done，使用 UseGlobal/DoneGlobal
	// app.UseGlobal(before)
	// app.DoneGlobal(after)
	app.Get("/main", mainHandler)
	// app.Get("/main", before, mainHandler, after)
	// 使用 ExecutionRules 去强制执行完成的处理程序，而不需要使用ctx.Next()
	// app.SetExecutionRules(iris.ExecutionRules{
	// 	// Begin: ...
	// 	// Main:  ...
	// 	Done: iris.ExecutionOptions{Force: true},
	// })

	app.Get("/", indexHandler).Name = "Index"

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

	// 方法: "POST"
	app.Post("/", handler)

	// 方法: "PUT"
	app.Put("/", handler)

	// 方法: "DELETE"
	app.Delete("/", handler)

	// 方法: "OPTIONS"
	app.Options("/", handler)

	// 方法: "TRACE"
	app.Trace("/", handler)

	// 方法: "CONNECT"
	app.Connect("/", handler)

	// 方法: "HEAD"
	app.Head("/", handler)

	// 方法: "PATCH"
	app.Patch("/", handler)

	// 注册支持所有 HTTP 方法的路由
	// app.Any("/", handler)

	none := app.None("/invisible/{username}", func(ctx iris.Context) {
		ctx.Writef("Hello %s with method: %s", ctx.Params().Get("username"), ctx.Method())

		if from := ctx.Values().GetString("from"); from != "" {
			ctx.Writef("\nI see that you're coming from %s", from)
		}
	})

	app.Get("/change", func(ctx iris.Context) {
		if none.IsOnline() {
			none.Method = iris.MethodNone
		} else {
			none.Method = iris.MethodGet
		}

		//刷新服务中重建的路由器，以便
		// 收到新路线通知。
		app.RefreshRouter()
	})

	app.Get("/execute", func(ctx iris.Context) {
		if !none.IsOnline() {
			ctx.Values().Set("from", "/execute with offline access")
			ctx.Exec("NONE", "/invisible/iris")
			return
		}

		// 与导航到 "http://localhost:8080/invisible/iris" 相同
		// 当 /change 被调用并且路由状态从
		// "离线" 改变为 "在线"
		ctx.Values().Set("from", "/execute")
		// 值和 session 可以被共享，
		// 当调用 Exec 从一个"外部"的 Context。
		// 	ctx.Exec("NONE", "/invisible/iris")
		// 或者在 "/change" 之后：
		ctx.Exec("GET", "/invisible/iris")
	})

	// 分组
	users := app.Party("/users", myMiddleware)
	// http://localhost:8080/users/42/profile
	users.Get("/{id:uint64}/profile", userProfileHandler)
	// http://localhost:8080/users/messages/1
	users.Get("/messages/{id:uint64}", userMessageHandler)

	app.PartyFunc("admins", func(admins iris.Party) {
		// http://localhost:8080/admins/42/profile
		admins.Get("/{id:uint64}/profile", userProfileHandler)
		// http://localhost:8080/admins/messages/1
		admins.Get("/messages/{id:uint64}", userMessageHandler)
	})

	// app.Listen(":8080")
	// 行为
	// app.Run(iris.Addr(":8080"), iris.WithoutPathCorrection) // 对请求的资源 禁用路径校正
	app.Run(iris.Addr(":8080"), iris.WithoutPathCorrectionRedirection) // 禁用路径校正和修正重定向
}

func before(ctx iris.Context) {
	info := "Welcome"
	requestPath := ctx.Path()
	println("Before the mainHandler: " + requestPath)

	ctx.Values().Set("info", info)
	ctx.Next()
}

func after(ctx iris.Context) {
	println("After the mainHandler")
}

func mainHandler(ctx iris.Context) {
	println("Inside mainHandler")

	info := ctx.Values().GetString("info")

	// 向客户端写一些内容作为响应。
	ctx.HTML("<h1>Response</h1>")
	ctx.HTML("<br/> Info: " + info)

	ctx.Next() // 执行 "after" 中间件
}

func indexHandler(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"message": "hello",
	})
}

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}

func handler(ctx iris.Context) {
	ctx.Writef("Hello from method: %s and path: %s\n", ctx.Method(), ctx.Path())
}

func userProfileHandler(ctx iris.Context) {
	ctx.Writef("user %s profile\n", ctx.Params().Get("id"))
}

func userMessageHandler(ctx iris.Context) {
	ctx.Writef("user %s messge\n", ctx.Params().Get("id"))
}
