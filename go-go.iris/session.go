package main

import (
	"github.com/kataras/iris/v12"

	"github.com/kataras/iris/v12/sessions"
)

var (
	cookieNameForSessionID = "mycookiesessionnameid"
	sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
	users                  = map[string]string{"test": "123456", "test2": "456789"}
)

func Secret(ctx iris.Context) {
	session := sess.Start(ctx)

	ctx.Application().Logger().Info(session.GetBoolean("logined"))
	ctx.Application().Logger().Info(session.GetBoolean("authenticated"))

	// 检查是否用户已经认证过
	if auth, _ := session.GetBoolean("authenticated"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}

	// 打印私密消息
	ctx.WriteString("Authorized!")
}

func Login(ctx iris.Context) {
	session := sess.Start(ctx)

	// 在这里验证身份
	user := ctx.PostValueDefault("user", "")
	pwd := ctx.PostValueDefault("pwd", "")

	if p, ok := users[user]; !ok || p != pwd {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}

	// auth := false
	// for u, p := range users {
	// 	if user == u && pwd == p {
	// 		auth = true
	// 		break
	// 	}
	// }
	// if !auth {
	// 	ctx.StatusCode(iris.StatusForbidden)
	// 	return
	// }

	// 设置一个认证过的用户
	session.Set("authenticated", true)
	session.Set("logined", true)
	ctx.WriteString("Logined")
}

func Logout(ctx iris.Context) {
	session := sess.Start(ctx)

	// 移除一个认证用户
	session.Set("authenticated", false)
	session.Set("logined", false)
	// 或者移除变量：
	session.Delete("authenticated")
	// 或者摧毁 session:
	session.Destroy()

	ctx.WriteString("logout")
}
