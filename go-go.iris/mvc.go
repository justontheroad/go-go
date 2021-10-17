package main

import (
	"github.com/kataras/iris/v12/mvc"
)

type ExampleController struct{}

func (c *ExampleController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/think/{idea:string}", "MyCustomHandler")
}

func (c *ExampleController) MyCustomHandler(idea string) string {
	return "MyCustomHandler says Hey" + idea
}

func (c *ExampleController) Get() mvc.Result {
	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>Welcome</h1>",
	}
}

func (c *ExampleController) GetPing() string {
	return "pong"
}

func (c *ExampleController) GetPingBy(name string) string {
	return "pong " + name
}
