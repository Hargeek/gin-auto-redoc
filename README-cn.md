[English](./README.md) | 简体中文

# gin-auto-redoc

用于自动为已包含 Swagger 文档的 Gin 应用程序注册 Redoc 文档的 Go 包

![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.22-%23007d9c)
![Gin Version](https://img.shields.io/badge/Gin-%3E%3D1.10-green)
[![GoDoc](https://godoc.org/github.com/hargeek/gin-auto-redoc?status.svg)](https://pkg.go.dev/github.com/hargeek/gin-auto-redoc)
[![Contributors](https://img.shields.io/github/contributors/hargeek/gin-auto-redoc)](https://github.com/hargeek/gin-auto-redoc/graphs/contributors)
[![License](https://img.shields.io/github/license/hargeek/gin-auto-redoc)](./LICENSE)

## 安装

```bash
go get github.com/hargeek/gin-auto-redoc
```

## 使用

在主文件中导入该包

```go
import "github.com/hargeek/gin-auto-redoc"
```

在 Gin 应用程序中注册 Redoc 中间件

```go
import (
    ginautodoc "github.com/hargeek/gin-auto-redoc"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.New()
    ginautodoc.Register(r)
    r.Run()
}
```

## 配置

默认情况下，Swagger 文档位于 /api/v1/swagger，Redoc 文档位于 /api/v1/doc。要自定义这些路径，请在注册中间件之前使用所需的配置调用 SetConfig

```go
ginautodoc.SetConfig(ginautodoc.Config{
    SwaggerPath: "/custom/swagger",
})
ginautodoc.Register(r)
```

## 访问 Redoc

启动 Gin 应用程序后，在 `http://localhost:<port>/api/v1/doc/index.html` 访问 Redoc 文档