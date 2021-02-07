# Env
- [Go mode](https://golang.org/ref/mod)
- [Learn](https://learn.go.dev/)
- [Go Packages](https://golang.org/pkg/)
- [tutorial](https://golang.org/doc/tutorial/getting-started)
- [Tour](https://pkg.go.dev/golang.org/x/tour)
- [External Packages](https://pkg.go.dev/)

```sh
# 安装Go
$ wget https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz
$ tar -C /usr/local -xzf go1.13.4.linux-amd64.tar.gz
$ mkdir -p /root/go
$ echo 'export GOROOT=/usr/local/go export PATH=$PATH:$GOROOT/bin export GOPATH=/root/go' >> /etc/profile
$ source /etc/profile

# 七牛源
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://goproxy.cn,direct
------
# 阿里源
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

# 13版本以下
$ export GOPROXY=https://mirrors.aliyun.com/goproxy/
```

```sh
# 初始化包管理器
$ go mod init demo
# 编译
$ go build
# 下载依赖
$ go get
# 查看项目依赖列表
$ go list -m all

$ go mod download 下载 go.mod 文件中指明的所有依赖

$ go mod tidy 整理现有的依赖，删除未使用的依赖。

$ go mod graph 查看现有的依赖结构

$ go mod init [module] module为项目目录名，生成 go.mod 文件 (Go 1.13 中唯一一个可以生成 go.mod 文件的子命令)

$ go mod edit 编辑 go.mod 文件

$ go mod vendor 导出现有的所有依赖 (事实上 Go modules 正在淡化 Vendor 的概念)

$ go mod verify 校验一个模块是否被篡改过

$ go clean -modcache 清理所有已缓存的模块版本数据。

$ go mod 查看所有 go mod的使用命令。
```

## Info

```go
// 方法定义
<-方法-><-方法名-><-参数名-><-参数类型-><-返回类型->
  func   Hello   (  name     string  )   string
// 说明
// 方法名 Hello 大写开头标识public 方法
// := 标识初始化并赋值
```

## Testing
```sh
# Add a file endwith [_test.go] in the same package which needs to test
## Sample : use hello_test.go to testing hello.go
$ go test
```

## Install
```shell
$ go build .
$ go install
# Run
$ study
```