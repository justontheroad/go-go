package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

// 创建你自己的自定义 Context，放入你需要的任何字段。
type CustomContext struct {
	// 嵌入 `iris.Context` -
	// 这是完全可选的， 但如果你需要
	// 不覆盖所有 Context 的方法！
	iris.Context
}

// 验证 CustomContext 实现 iris.Context 在编译的时候。
var _ iris.Context = &CustomContext{}

func (ctx *CustomContext) Do(handlers context.Handlers) {
	context.Do(ctx, handlers)
}

func (ctx *CustomContext) Next() {
	context.Next(ctx)
}

// 处覆盖想要的任何 Context 方法
func (ctx *CustomContext) HTML(format string, args ...interface{}) (int, error) {
	ctx.Application().Logger().Infof("Executing .HTML function from CustomContext")

	ctx.ContentType("text/html")
	return ctx.Writef(format, args...)
}
