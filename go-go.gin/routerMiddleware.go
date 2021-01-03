package main

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 自定义路由中间件，*gin.Content中定义了example变量
func CustomRouterMiddleware(c *gin.Context) {
	t := time.Now()
	log.Print("中间件开始执行了")
	// 在gin上下文中定义一个变量
	c.Set("example", "CustomRouterMiddle1")

	// 请求之前

	// c.Next()
	c.Abort()

	// 请求之后
	latency := time.Since(t)
	log.Print(latency)

	status := c.Writer.Status()
	log.Println(status)
}

func FirstMiddleware(c *gin.Context) {
	log.Print("first middleware before next()")
	isAbort := c.Query("isAbort")
	bAbort, err := strconv.ParseBool(isAbort)
	if err != nil {
		log.Printf("is abort value err, value %s\n", isAbort)
		c.Next() // (2)
	}
	if bAbort {
		log.Print("first middleware abort") //(3)
		c.Abort()
		//c.AbortWithStatusJSON(http.StatusOK, "abort is true")
		return
	} else {
		log.Print("first middleware doesnot abort") //(4)
		return
	}

	log.Print("first middleware after next()")
}

func SecondMiddleware(c *gin.Context) {
	log.Print("current inside of second middleware")
}
