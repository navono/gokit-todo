# go-kit-demo
演示`Go-Kit`的示例代码。

# dev Environment
下载包时，需要科学上网，同时设置`http_proxy`和`https_proxy`。比如在`PowerShell`中：
> $env:http_proxy="http://127.0.0.1:1080"
<br>
> $env:https_proxy="http://127.0.0.1:1080"

然后再运行以下命令。

## 下载`GoKit CLI`
这个是一个命令行工具，用来产生`go-kit`可用的模板代码：
> go get github.com/kujtimiihoxha/kit

## 安装
> go install github.com/kujtimiihoxha/kit

## 依赖
> dep ensure -add -v github.com/go-kit/kit

# 创建服务
## 生成服务模板
> kit new service todo

## 定义 endpoint
```
pkg
  io
    io.go
```
使用`mongodb`来存储。然后在 `service.go` 中定义`endpoints`的接口：
```go
// TodoService describes the service.
type TodoService interface {
	Get(ctx context.Context) (t []io.Todo, error error)
	Add(ctx context.Context, todo io.Todo) (t io.Todo, error error)
	SetComplete(ctx context.Context, id string) (error error)
	RemoveComplete(ctx context.Context, id string) (error error)
	Delete(ctx context.Context, id string) (error error)
}
```

# 生成服务代码
使用`Gorilla`作为默认的`http` handler。
> kit g s todo -w --gorilla

安装依赖：
> dep ensure

# 运行
> go run todo/cmd/main.go

如果出现以下错误：
> #github.com/navono/go-kit-demo/vendor/github.com/openzipkin/zipkin-go-opentracing/thrift/gen-go/scribe
> vendor\github.com\openzipkin\zipkin-go-opentracing\thrift\gen-go\scribe\scribe.go:283:14: too many arguments in call to oprot.Flush

在`Gopkg.toml`中，增加以下字段：
>[[override]]
  name = "github.com/apache/thrift"
  version = "master"

# 存储（DB）
