package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// 禁用控制台颜色
	gin.DisableConsoleColor()
	// 使用默认中间件创建一个gin路由器
	// logger and recovery (crash-free) 中间件
	// r := gin.Default()
	// 无中间件启动
	r := gin.New()
	// 使用Logger中间件
	// r.Use(gin.Logger())
	// 使用自定义格式的Logger中间件
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	// 使用Recovery中间件
	r.Use(gin.Recovery())

	// 创建日志
	f, _ := os.Create("run.log")
	gin.DefaultWriter = io.MultiWriter(f)
	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 获取路径中的参数
	// 此规则能够匹配/user/{name}这种格式
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 路由模糊匹配
	// 此规则能够匹配/user/{name}/这种格式，也能匹配/user/{name}/{*}
	r.GET("/user/:name/*action", func(c *gin.Context) {
		// 获取API参数
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// GET获取请求参数
	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // 是 c.Request.URL.Query().Get("lastname") 的简写
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	// POST获取请求参数
	r.POST("user/:name/info", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous") // 此方法可以设置默认值
		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	// 上传单个文件
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			return
		}
		log.Println(file.Filename)

		c.SaveUploadedFile(file, "./upload/"+file.Filename)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	// 表单限制上传大小 (默认 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	// 上传多个文件
	r.POST("/multi-upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			return
		}
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)

			// 上传文件到指定的路径
			c.SaveUploadedFile(file, "./upload/"+file.Filename)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	// 路由组
	v1 := r.Group("v1")
	{
		v1.GET("/hello", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello")
		})
		v1.GET("/world", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello world")
		})
	}
	v2 := r.Group("v2")
	{
		v2.GET("/hello", func(c *gin.Context) {
			c.String(http.StatusOK, "V2 Hello")
		})
		v2.GET("/world", func(c *gin.Context) {
			c.String(http.StatusOK, "V2 Hello world")
		})
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
	// router.Run(":8080") // 指定监听地址

	// router := NewRouter()
	// log.Fatal(http.ListenAndServe(":9090", router))
}
