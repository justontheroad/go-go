package main

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/core/errgroup"
	"github.com/kataras/iris/v12/sessions"
)

var (
	g errgroup.Error
)

func main() {
	// Creates a iris app with default middleware:
	// logger and recovery (crash-free) middleware
	app := iris.Default()
	// custom middleware
	app.Use(MyMiddleware)
	// Use 和 Done 应用于当前路由分组和它的子分组，即在调用 Use 或 Done 之前的路由，不会应用该中间件
	app.Get("/home", IndexHandler)
	app.Use(Before)
	app.Done(After)
	// 在路由注册前使用 `app.Use/Done，使用 UseGlobal/DoneGlobal
	// app.UseGlobal(before)
	// app.DoneGlobal(after)
	app.Get("/main", MainHandler)
	// app.Get("/main", before, mainHandler, after)
	// 使用 ExecutionRules 去强制执行完成的处理程序，而不需要使用ctx.Next()
	// app.SetExecutionRules(iris.ExecutionRules{
	// 	// Begin: ...
	// 	// Main:  ...
	// 	Done: iris.ExecutionOptions{Force: true},
	// })

	app.Get("/", IndexHandler).Name = "Index"

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

	// app.Get("{root:path}", func(ctx iris.Context) {
	// 	root := ctx.Params().Get("root")
	// 	ctx.JSON(iris.Map{
	// 		"root": root,
	// 	})
	// })

	// 方法: "POST"
	app.Post("/", CustomHandler)

	// 方法: "PUT"
	app.Put("/", CustomHandler)

	// 方法: "DELETE"
	app.Delete("/", CustomHandler)

	// 方法: "OPTIONS"
	app.Options("/", CustomHandler)

	// 方法: "TRACE"
	app.Trace("/", CustomHandler)

	// 方法: "CONNECT"
	app.Connect("/", CustomHandler)

	// 方法: "HEAD"
	app.Head("/", CustomHandler)

	// 方法: "PATCH"
	app.Patch("/", CustomHandler)

	// 注册支持所有 HTTP 方法的路由
	// app.Any("/", CustomHandler)

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
	users := app.Party("/users", MyMiddleware)
	// http://localhost:8080/users/42/profile
	users.Get("/{id:uint64}/profile", UserProfileHandler)
	// http://localhost:8080/users/messages/1
	users.Get("/messages/{id:uint64}", UserMessageHandler)

	app.PartyFunc("admins", func(admins iris.Party) {
		// http://localhost:8080/admins/42/profile
		admins.Get("/{id:uint64}/profile", UserProfileHandler)
		// http://localhost:8080/admins/messages/1
		admins.Get("/messages/{id:uint64}", UserMessageHandler)
	})

	// 注册 view
	app.RegisterView(iris.HTML("./views", ".html"))

	// 错误处理
	app.OnErrorCode(iris.StatusNotFound, NotFound)
	app.OnErrorCode(iris.StatusInternalServerError, InternalServerError)
	// 为所有“错误”注册一个处理程序
	// 状态代码(kataras/iris/context.StatusCodeNotSuccessful)
	// app.OnAnyErrorCode(errorHandler)

	// 用本地 net/http 处理程序包装路由器。
	app.WrapRouter(func(w http.ResponseWriter, r *http.Request, router http.HandlerFunc) {
		path := r.URL.Path
		// 判断路由前缀
		if strings.HasPrefix(path, "/other") {
			// 获取并释放上下文以便使用它来执行
			ctx := app.ContextPool.Acquire(w, r)
			OtherHandler(ctx)
			app.ContextPool.Release(ctx)
			return
		}

		// 否则继续照常服务路由。
		router.ServeHTTP(w, r)
	})

	// app.Listen(":8080")
	// 行为
	// app.Run(iris.Addr(":8080"), iris.WithoutPathCorrection) // 对请求的资源 禁用路径校正
	// app.Run(iris.Addr(":8080"), iris.WithoutPathCorrectionRedirection) // 禁用路径校正和修正重定向

	customApp := iris.New()
	customApp.ContextPool.Attach(func() iris.Context {
		return &CustomContext{
			// 如果你要使用嵌入式 Context,
			// 调用 `context.NewContext` 创建一个：
			Context: context.NewContext(customApp),
		}
	})

	//  在 ./view/** 目录中的 .html 文件上注册视图引擎
	customApp.RegisterView(iris.HTML("./views", ".html"))

	customApp.Handle("GET", "/", recordWhichContextForExample,
		func(ctx iris.Context) {
			// 使用 覆盖过的 Context 的 HTML 方法。
			ctx.HTML("<h1> Hello from my custom context's HTML! </h1>")
		})

	customApp.Handle("GET", "/hi/{firstname:alphabetical}", recordWhichContextForExample,
		func(ctx iris.Context) {
			// firstname := ctx.Values().GetString("firstname")
			firstname := ctx.Params().GetString("firstname")
			ctx.ViewData("firstname", firstname)
			ctx.Gzip(true)

			ctx.View("hi.html")
		})

	// 当执行 control+C/cmd+C  时关闭连接
	iris.RegisterOnInterrupt(func() {
		sessRedisDb.Close()
	})

	// 3. session 注册 reids db支持
	sess.UseDatabase(sessRedisDb)
	app.Use(sess.Handler())
	app.Get("/secret", Secret)
	app.Post("/login", Login, func(ctx iris.Context) {
		session := sessions.Get(ctx)
		ctx.Application().Logger().Info(session.GetBoolean("authenticated"))
	})
	app.Post("/logout", Logout)

	// listen 阻塞代码，多端口监听需要使用 coroutine
	go app.Run(iris.Addr(":8080"), iris.WithoutPathCorrectionRedirection) // 禁用路径校正和修正重定向
	customApp.Listen(":8888")
}

func recordWhichContextForExample(ctx iris.Context) {
	ctx.Application().Logger().Infof("(%s) Handler is executing from: '%s'",
		ctx.Path(), reflect.TypeOf(ctx).Elem().Name())

	ctx.Next()
}
