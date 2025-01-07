[English](./README.md) | 简体中文

# gin-auto-redoc

用于自动为 Gin 应用程序注册 Redoc 文档的 Go 包

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
    _ "github.com/hargeek/gin-auto-redoc"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.New()
    ginautodoc.Register(r)
    r.Run()
}
```

## 配置

可以在注册中间件之前调用 SetConfig 来设置自定义配置

```go
ginautodoc.SetConfig(ginautodoc.Config{
    SwaggerPath: "/custom/swagger",
})
ginautodoc.Register(r)
```