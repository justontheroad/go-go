1. brew 安装
```
brew search go
go@1.10                                  hugo
go@1.11                                  jpegoptim
go@1.12                                  katago
go@1.13                                  lego
go@1.9                                   lgogdownloader
brew install go@1.13
```
2. 下载golang安装包，根据提示安装。
[下载地址](https://studygolang.com/dl/golang/go1.15.4.darwin-amd64.pkg)
```
// 安装目录
/usr/local/go
total 376
-rw-r--r--    1 root  wheel  55669 11  6 05:24 AUTHORS
-rw-r--r--    1 root  wheel   1339 11  6 05:24 CONTRIBUTING.md
-rw-r--r--    1 root  wheel  95475 11  6 05:24 CONTRIBUTORS
-rw-r--r--    1 root  wheel   1479 11  6 05:24 LICENSE
-rw-r--r--    1 root  wheel   1303 11  6 05:24 PATENTS
-rw-r--r--    1 root  wheel   1607 11  6 05:24 README.md
-rw-r--r--    1 root  wheel    397 11  6 05:24 SECURITY.md
-rw-r--r--    1 root  wheel      8 11  6 05:24 VERSION
drwxr-xr-x   21 root  wheel    672 11  6 05:24 api
drwxr-xr-x    4 root  wheel    128 11  6 05:50 bin
drwxr-xr-x   45 root  wheel   1440 11  6 05:24 doc
-rw-r--r--    1 root  wheel   5686 11  6 05:24 favicon.ico
drwxr-xr-x    3 root  wheel     96 11  6 05:24 lib
drwxr-xr-x   14 root  wheel    448 11  6 05:24 misc
drwxr-xr-x    6 root  wheel    192 11  6 05:27 pkg
-rw-r--r--    1 root  wheel     26 11  6 05:24 robots.txt
drwxr-xr-x   69 root  wheel   2208 11  6 05:24 src
drwxr-xr-x  328 root  wheel  10496 11  6 05:24 test
```
3. 配置环境变量，vi ~/.bashrc，添加如下内容，source ~/.bashrc
```
# GOROOT
export GOROOT=/usr/local/go

# GOPATH，GOPATH可以根据个人习惯设置为其他目录
export GOPATH=$HOME/Documents/code/gopath

# GOPATH root bin
export PATH=$PATH:$GOROOT/bin
```
- source命令作用#
在当前bash环境下读取并执行FileName中的命令
4. 验证go是否安装成功
```
go version
    go version go1.15.4 darwin/amd64
go env
    GO111MODULE=""
    GOARCH="amd64"
    GOBIN=""
    GOCACHE="/Users/ontheway/Library/Caches/go-build"
    GOENV="/Users/ontheway/Library/Application Support/go/env"
    GOEXE=""
    GOFLAGS=""
    GOHOSTARCH="amd64"
    GOHOSTOS="darwin"
    GOINSECURE=""
    GOMODCACHE="$HOME/Documents/code/gopath/pkg/mod"
    GONOPROXY=""
    GONOSUMDB=""
    GOOS="darwin"
    GOPATH="$HOME/Documents/code/gopath/"
    GOPRIVATE=""
    GOPROXY="https://proxy.golang.org,direct"
    GOROOT="/usr/local/go"
    GOSUMDB="sum.golang.org"
    GOTMPDIR=""
    GOTOOLDIR="/usr/local/go/pkg/tool/darwin_amd64"
    GCCGO="gccgo"
    AR="ar"
    CC="clang"
    CXX="clang++"
    CGO_ENABLED="1"
    GOMOD=""
    CGO_CFLAGS="-g -O2"
    CGO_CPPFLAGS=""
    CGO_CXXFLAGS="-g -O2"
    CGO_FFLAGS="-g -O2"
    CGO_LDFLAGS="-g -O2"
    PKG_CONFIG="pkg-config"
    GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/qd/x1xjd5sd0cg4mxy5qftchclw0000gn/T/go-build325734103=/tmp/go-build -gno-record-gcc-switches -fno-common"
```
5. 设置git账户
```
git config --global user.name "your name"
git config --global user.email "your mail"
```
6. git 配置ssh key
```
## 生成密钥
ssh-keygen -t rsa -C "your mail"
## 确认秘钥的保存路径（如果不需要改路径则直接回车）
Generating public/private rsa key pair.
Enter file in which to save the key (/Users/ontheway/.ssh/id_rsa): 
## 创建密码（如果不需要密码则直接回车）
Enter passphrase (empty for no passphrase):
## 确认密码；
Enter same passphrase again: 
Your identification has been saved in /Users/ontheway/.ssh/id_rsa.
Your public key has been saved in /Users/ontheway/.ssh/id_rsa.pub.
The key fingerprint is:
SHA256:ufLIr0Q8uIkRPSrAZi749vUJolSPU7yqoZNElaIIxtk 490421206@qq.com
The key's randomart image is:
+---[RSA 2048]----+
|. o .            |
|o= E             |
|*++ o            |
|O. o =   .       |
|=.o o * S        |
|.+ + B o .       |
|..* * B .        |
|o+ + B * .       |
|..o.o +o=        |
+----[SHA256]-----+
```
- github配置shh，选择SSH and GPG keys项，new SSH keys，将上一步骤生成的公钥输入并保存
    - 默认文件地址：/Users/{user}/.ssh/id_rsa.pub
    - 测试git clone
    ```
    git clone git@github.com:justontheroad/go-go.git
    Cloning into 'go-go'...
    The authenticity of host 'github.com (13.229.188.59)' can't be established.
    RSA key fingerprint is SHA256:nThbg6kXUpJWGl7E1IGOCspRomTxdCARLviKw6E5SY8.
    Are you sure you want to continue connecting (yes/no)?
    remote: Enumerating objects: 3, done.
    remote: Counting objects: 100% (3/3), done.
    Receiving objects: 100% (3/3), done.
    remote: Total 3 (delta 0), reused 0 (delta 0), pack-reused 0
    ```
7. IDE，使用vscode + go插件，也可选择上goland

> [Download and install](https://docs.studygolang.com/doc/install)