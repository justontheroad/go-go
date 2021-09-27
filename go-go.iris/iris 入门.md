### Iris Web Framework
Iris 是一个用 Go 编写的快速，简单但功能齐全且非常高效的 Web 框架。官网：https://www.iris-go.com/

1. 安装 Iris
    1. 使用go get 安装
    ```
    go get github.com/kataras/iris/v12@latest
    ```
    2. go moudle 安装
    ```
    // model 初始化，go-go/go-go.iris为项目名
    go mod init go-go/go-go.iris
    ```
    ```
    module go-go/go-go.iris
    go 1.17
    // 文件中 import iris package
    require (
        github.com/kataras/iris/v12 latest
    )
    ```
    ```
    // 检测依赖
    go mod tidy
    // 下载依赖
    go mod download
    // 导入依赖到本地vendor目录，根据实际需要确认是否要执行
    go mod vendor
    ```
2. 代码中导入
    ```
    import "github.com/kataras/iris/v12"
    ```
3. 快速入门，demo
    ```
    package main

    import "github.com/kataras/iris/v12"
    
    func main() {
        app := iris.New()
        app.Get("/", index)
        app.Listen(":8080")
        // app.Run(iris.Addr(":8080"))
    }

    func index(ctx iris.Context) {
        ctx.JSON(iris.Map{
            "message": "hello",
        })
    }
    ```
    - 通过 http://localhost:8080 或 http://127.0.0.1:8080 访问
4. 使用 GET, POST, PUT, PATCH, DELETE, OPTIONS
    ```
    func main() {
    	// Creates a iris app with default middleware:
    	// logger and recovery (crash-free) middleware
    	app := iris.Default()
    
    	app.Get("/someGet", handler)
    	app.Post("/somePost", handler)
    	app.Put("/somePut", handler)
    	app.Delete("/someDelete", handler)
    	app.Patch("/somePatch", handler)
    	app.Head("/someHead", handler)
    	app.Options("/someOptions", handler)
        app.Trace("/", handler)
        app.Connect("/", handler)
        // 注册支持所有 HTTP 方法的路由
        app.Any("/", handler)
    
    	app.Run(iris.Addr(":8080"))
    }

    func handler(ctx iris.Context){
        ctx.Writef("Hello from method: %s and path: %s\n", ctx.Method(), ctx.Path())
    }
    ```
5. 获取路径中的参数
    ```
    app := iris.Default()
    // 匹配所有以 "/user/" 为前缀的 GET 请求，后跟随单个路径部分，:string为类型限定
    app.Get("/user/{name:string}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.JSON(iris.Map{
			"name": name,
		})
	})
	// 匹配所有包含 "/assets/**/*" 前缀的 GET 请求， 它是一个 ctx.Params().Get("asset") 的通配符，等同于任何跟随在 /assets/之后的路径
	app.Get("/assets/{asset:path}", func(ctx iris.Context) {
		ctx.JSON(ctx.Params().Get("asset"))
	})
    // 匹配所有以 /user/ 为前缀的GET请求，后跟一个等于或大于 1 的数字
    app.Get("/user/{userid:int min(1)}", func(ctx iris.Context) {
		id := ctx.Params().Get("userid")
		ctx.JSON(iris.Map{
			"id": id,
		})
	})
    // 匹配除其他路由已处理的请求之外的所有 GET 请求
    app.Get("{root:path}", func(ctx iris.Context) {
        root := ctx.Params().Get("root")
		ctx.JSON(iris.Map{
			"root": root,
		})
    })
    ```
6. 获取Get参数、Post参数
    ```
    ```
7. 路由分组
    ```

    ```

> [gin github](https://github.com/gin-gonic/gin)