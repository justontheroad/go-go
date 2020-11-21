### Module
Go 1.11 开始，Go 允许在 $GOPATH/src 外的任何目录下使用 go.mod 创建项目。在 $GOPATH/src 中，为了兼容性，Go 命令仍然在旧的 GOPATH 模式下运行。从 Go 1.13 开始，模块模式将成为默认模式。

- 解决go被墙问题
```
go: finding module for package github.com/gin-gonic/gin
goModules.go:4:2: module github.com/gin-gonic/gin: Get "https://proxy.golang.org/github.com/gin-gonic/gin/@v/list": dial tcp 172.217.160.113:443: i/o timeout
```


> [Go Modules 详解使用](https://learnku.com/articles/27401)