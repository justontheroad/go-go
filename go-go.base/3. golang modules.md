### Module
Go 1.11 开始，Go 允许在 $GOPATH/src 外的任何目录下使用 go.mod 创建项目。在 $GOPATH/src 中，为了兼容性，Go 命令仍然在旧的 GOPATH 模式下运行。从 Go 1.13 开始，模块模式将成为默认模式。

- 解决proxy.golang.org被墙问题，需要使用国内代理，执行命令 go env -w GOPROXY=https://goproxy.cn
```
go: finding module for package github.com/gin-gonic/gin
goModules.go:4:2: module github.com/gin-gonic/gin: Get "https://proxy.golang.org/github.com/gin-gonic/gin/@v/list": dial tcp 172.217.160.113:443: i/o timeout
```
- GO111MODULE，golang版本管理工具
    - 用环境变量 GO111MODULE 开启或关闭模块支持，它有三个可选值：off、on、auto，默认值是 auto。
    1. GO111MODULE=off 无模块支持，go 会从 GOPATH 和 vendor 文件夹寻找包。
    2. GO111MODULE=on 模块支持，go 会忽略 GOPATH 和 vendor 文件夹，只根据 go.mod 下载依赖。
    3. GO111MODULE=auto 在 $GOPATH/src 外面且根目录有 go.mod 文件时，开启模块支持。
    - 在使用模块的时候，GOPATH 是无意义的，不过它还是会把下载的依赖储存在 $GOPATH/src/mod 中，也会把 go install 的结果放在 $GOPATH/bin 中。
1. 创建一个新模块
    1. 可以在 $GOPATH/src 之外的任何地方创建一个新的目录，例如在项目中创建 src，并初始化
    ```
    mkdir src && cd src
    // 初始化
    go mod init src
    ```
    2. 初始化成功之后，目录下会生成一个 go.mod 文件
    ```
    module src

    go 1.15
    ```
2. 添加依赖项
    1. 测试项目中使用gin作为例子，在main.go中import gin 包
    ```
    package main

    import "github.com/gin-gonic/gin"
    
    func main() {
        r := gin.Default()
        r.GET("/ping", func(c *gin.Context) {
            c.JSON(200, gin.H{
                "message": "pong",
            })
        })
        r.Run() // listen and serve on 0.0.0.0:8080
    }
    ```
    2. go build，命令执行完成之后，go.mod被修改，并新增了go.sum文件。
        - go.mod 则是描述直接依赖包(相当于glide.yaml)
        - go.sum 则是描述依赖树锁定(相当于glide.lock)
    ```
    module src

    go 1.15
    
    require github.com/gin-gonic/gin v1.6.3
    ```
    ```
    github.com/davecgh/go-spew v1.1.0/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
    github.com/davecgh/go-spew v1.1.1/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
    github.com/gin-contrib/sse v0.1.0 h1:Y/yl/+YNO8GZSjAhjMsSuLt29uWRFHdHYUb5lYOV9qE=
    github.com/gin-contrib/sse v0.1.0/go.mod h1:RHrZQHXnP2xjPF+u1gW/2HnVO7nvIa9PG3Gm+fLHvGI=
    github.com/gin-gonic/gin v1.1.4 h1:XLaCFbU39SSGRQrEeP7Z7mM3lvRqC4vE5tEaVdLDdSE=
    github.com/gin-gonic/gin v1.1.4/go.mod h1:7cKuhb5qV2ggCFctp2fJQ+ErvciLZrIeoOSOm6mUr7Y=
    github.com/gin-gonic/gin v1.6.3 h1:ahKqKTFpO5KTPHxWZjEdPScmYaGtLo8Y4DMHoEsnp14=
    github.com/gin-gonic/gin v1.6.3/go.mod h1:75u5sXoLsGZoRN5Sgbi1eraJ4GU3++wFwWzhwvtwp4M=
    github.com/go-playground/assert/v2 v2.0.1/go.mod h1:VDjEfimB/XKnb+ZQfWdccd7VUvScMdVu0Titje2rxJ4=
    github.com/go-playground/locales v0.13.0 h1:HyWk6mgj5qFqCT5fjGBuRArbVDfE4hi8+e8ceBS/t7Q=
    github.com/go-playground/locales v0.13.0/go.mod h1:taPMhCMXrRLJO55olJkUXHZBHCxTMfnGwq/HNwmWNS8=
    github.com/go-playground/universal-translator v0.17.0 h1:icxd5fm+REJzpZx7ZfpaD876Lmtgy7VtROAbHHXk8no=
    github.com/go-playground/universal-translator v0.17.0/go.mod h1:UkSxE5sNxxRwHyU+Scu5vgOQjsIJAF8j9muTVoKLVtA=
    github.com/go-playground/validator/v10 v10.2.0 h1:KgJ0snyC2R9VXYN2rneOtQcw5aHQB1Vv0sFl1UcHBOY=
    github.com/go-playground/validator/v10 v10.2.0/go.mod h1:uOYAAleCW8F/7oMFd6aG0GOhaH6EGOAJShg8Id5JGkI=
    github.com/golang/protobuf v1.3.3 h1:gyjaxf+svBWX08ZjK86iN9geUJF0H6gp2IRKX6Nf6/I=
    github.com/golang/protobuf v1.3.3/go.mod h1:vzj43D7+SQXF/4pzW/hwtAqwc6iTitCiVSaWz5lYuqw=
    github.com/google/gofuzz v1.0.0/go.mod h1:dBl0BpW6vV/+mYPU4Po3pmUjxk6FQPldtuIdl/M65Eg=
    github.com/json-iterator/go v1.1.9 h1:9yzud/Ht36ygwatGx56VwCZtlI/2AD15T1X2sjSuGns=
    github.com/json-iterator/go v1.1.9/go.mod h1:KdQUCv79m/52Kvf8AW2vK1V8akMuk1QjK/uOdHXbAo4=
    github.com/leodido/go-urn v1.2.0 h1:hpXL4XnriNwQ/ABnpepYM/1vCLWNDfUNts8dX3xTG6Y=
    github.com/leodido/go-urn v1.2.0/go.mod h1:+8+nEpDfqqsY+g338gtMEUOtuK+4dEMhiQEgxpxOKII=
    github.com/manucorporat/sse v0.0.0-20160126180136-ee05b128a739 h1:ykXz+pRRTibcSjG1yRhpdSHInF8yZY/mfn+Rz2Nd1rE=
    github.com/manucorporat/sse v0.0.0-20160126180136-ee05b128a739/go.mod h1:zUx1mhth20V3VKgL5jbd1BSQcW4Fy6Qs4PZvQwRFwzM=
    github.com/mattn/go-isatty v0.0.12 h1:wuysRhFDzyxgEmMf5xjvJ2M9dZoWAXNNr5LSBS7uHXY=
    github.com/mattn/go-isatty v0.0.12/go.mod h1:cbi8OIDigv2wuxKPP5vlRcQ1OAZbq2CE4Kysco4FUpU=
    github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 h1:ZqeYNhU3OHLH3mGKHDcjJRFFRrJa6eAM5H+CtDdOsPc=
    github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421/go.mod h1:6dJC0mAP4ikYIbvyc7fijjWJddQyLn8Ig3JB5CqoB9Q=
    github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742 h1:Esafd1046DLDQ0W1YjYsBW+p8U2u7vzgW2SQVmlNazg=
    github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742/go.mod h1:bx2lNnkwVCuqBIxFjflWJWanXIb3RllmbCylyMrvgv0=
    github.com/pmezard/go-difflib v1.0.0/go.mod h1:iKH77koFhYxTK1pcRnkKkqfTogsbg7gZNVY4sRDYZ/4=
    github.com/stretchr/objx v0.1.0/go.mod h1:HFkY916IF+rwdDfMAkV7OtwuqBVzrE8GR6GFx+wExME=
    github.com/stretchr/testify v1.3.0/go.mod h1:M5WIy9Dh21IEIfnGCwXGc5bZfKNJtfHm1UVUgZn+9EI=
    github.com/stretchr/testify v1.4.0/go.mod h1:j7eGeouHqKxXV5pUuKE4zz7dFj8WfuZ+81PSLYec5m4=
    github.com/ugorji/go v1.1.7 h1:/68gy2h+1mWMrwZFeD1kQialdSzAb432dtpeJ42ovdo=
    github.com/ugorji/go v1.1.7/go.mod h1:kZn38zHttfInRq0xu/PH0az30d+z6vm202qpg1oXVMw=
    github.com/ugorji/go/codec v1.1.7 h1:2SvQaVZ1ouYrrKKwoSk2pzd4A9evlKJb9oTL+OaLUSs=
    github.com/ugorji/go/codec v1.1.7/go.mod h1:Ax+UKWsSmolVDwsd+7N3ZtXu+yMGCf907BLYF3GoBXY=
    golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2/go.mod h1:djNgcEr1/C05ACkg1iLfiJU5Ep61QUkGW8qpdssI0+w=
    golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9/go.mod h1:LzIPMQfyMNhhGPhUkYOs5KpL4U8rLKemX1yGLhDgUto=
    golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3/go.mod h1:t9HGtf8HONx5eT2rtn7q6eTqICYqUVnKs3thJo3Qplg=
    golang.org/x/net v0.0.0-20201110031124-69a78807bb2b h1:uwuIcX0g4Yl1NC5XAz37xsr2lTtcqevgzYNVt49waME=
    golang.org/x/net v0.0.0-20201110031124-69a78807bb2b/go.mod h1:sp8m0HH+o8qH0wwXwYZr8TS3Oi6o0r6Gce1SSxlDquU=
    golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a/go.mod h1:STP8DvDyc/dI5b8T5hshtkjS+E42TnysNCUPdjciGhY=
    golang.org/x/sys v0.0.0-20190412213103-97732733099d/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
    golang.org/x/sys v0.0.0-20200116001909-b77594299b42 h1:vEOn+mP2zCOVzKckCZy6YsCtDblrpj/w7B9nxGNELpg=
    golang.org/x/sys v0.0.0-20200116001909-b77594299b42/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
    golang.org/x/sys v0.0.0-20200930185726-fdedc70b468f h1:+Nyd8tzPX9R7BWHguqsrbFdRx3WQ/1ib8I44HXV5yTA=
    golang.org/x/sys v0.0.0-20200930185726-fdedc70b468f/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
    golang.org/x/text v0.3.0/go.mod h1:NqM8EUOU14njkJ3fqMW+pc6Ldnwhi/IjpwHt7yyuwOQ=
    golang.org/x/text v0.3.2/go.mod h1:bEr9sfX3Q8Zfm5fL9x+3itogRgK3+ptLWKqgva+5dAk=
    golang.org/x/text v0.3.3/go.mod h1:5Zoc/QRtKVWzQhOtBMvqHzDpF6irO9z98xDceosuGiQ=
    golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e/go.mod h1:n7NCudcB/nEzxVGmLbDWY5pfWTLqBcC2KZ6jyYvM4mQ=
    gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
    gopkg.in/go-playground/validator.v8 v8.18.2 h1:lFB4DoMU6B626w8ny76MV7VX6W2VHct2GVOI3xgiMrQ=
    gopkg.in/go-playground/validator.v8 v8.18.2/go.mod h1:RX2a/7Ha8BgOhfk7j780h4/u/RRjR0eouCJSH80/M2Y=
    gopkg.in/yaml.v2 v2.2.2/go.mod h1:hI93XBmqTisBFMUTm0b8Fm+jr3Dg1NNxqwp+5A1VGuI=
    gopkg.in/yaml.v2 v2.2.8 h1:obN1ZagJSUGI0Ek/LBmuj4SNLPfIny3KsKFopxRdj10=
    gopkg.in/yaml.v2 v2.2.8/go.mod h1:hI93XBmqTisBFMUTm0b8Fm+jr3Dg1NNxqwp+5A1VGuI=

    ```
    - go mod 还支持以下命令
    ```
    he commands are:
    download    download modules to local cache
    edit        edit go.mod from tools or scripts
    graph       print module requirement graph
    init        initialize new module in current directory
    tidy        add missing and remove unused modules
    vendor      make vendored copy of dependencies
    verify      verify dependencies have expected content
    why         explain why packages or modules are needed
    ```
3. 升级依赖项
    1. 查看项目使用的依赖包
    ```
    $ go list -m all
    src
    github.com/davecgh/go-spew v1.1.1
    github.com/gin-contrib/sse v0.1.0
    github.com/gin-gonic/gin v1.6.3
    github.com/go-playground/assert/v2 v2.0.1
    github.com/go-playground/locales v0.13.0
    github.com/go-playground/universal-translator v0.17.0
    github.com/go-playground/validator/v10 v10.2.0
    github.com/golang/protobuf v1.3.3
    github.com/google/gofuzz v1.0.0
    github.com/json-iterator/go v1.1.9
    github.com/leodido/go-urn v1.2.0
    github.com/mattn/go-isatty v0.0.12
    github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421
    github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742
    github.com/pmezard/go-difflib v1.0.0
    github.com/stretchr/objx v0.1.0
    github.com/stretchr/testify v1.4.0
    github.com/ugorji/go v1.1.7
    github.com/ugorji/go/codec v1.1.7
    golang.org/x/sys v0.0.0-20200116001909-b77594299b42
    golang.org/x/text v0.3.2
    golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e
    gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405
    gopkg.in/yaml.v2 v2.2.8
    ```
    2. 查看依赖的版本历史
    ```
    $ go list -m -versions github.com/gin-gonic/gin
    github.com/gin-gonic/gin v1.1.1 v1.1.2 v1.1.3 v1.1.4 v1.3.0 v1.4.0 v1.5.0 v1.6.0 v1.6.1 v1.6.2 v1.6.3
    ```
    3. 回退到历史版本
    ```
    ## 使用ge get命令
    $ go get github.com/gin-gonic/gin@v1.6.2
    ## 或者修改mod
    $ go mod edit -require="github.com/gin-gonic/gin@v1.6.2"
    $ go mod tidy // 下载更新依赖
    go: downloading github.com/gin-gonic/gin v1.6.2
    go: downloading github.com/go-playground/assert/v2 v2.0.1
    ```
4. 删除未使用的依赖项
    - go.tidy 会自动清理掉不需要的依赖项，同时可以将依赖项更新到当前版本
    ```
    $ go mod tidy
    ```

> [Go Modules 详解使用](https://learnku.com/articles/27401)
> [golang版本管理工具GO111MODULE](https://www.cnblogs.com/embedded-linux/p/11616183.html)