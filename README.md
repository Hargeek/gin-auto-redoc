[简体中文](./README-cn.md) | English

# gin-auto-redoc

A Go package to automatically register Redoc documentation for Gin applications

![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.22-%23007d9c)
[![GoDoc](https://godoc.org/github.com/hargeek/gin-auto-redoc?status.svg)](https://pkg.go.dev/github.com/hargeek/gin-auto-redoc)
[![Contributors](https://img.shields.io/github/contributors/hargeek/gin-auto-redoc)](https://github.com/hargeek/gin-auto-redoc/graphs/contributors)
[![License](https://img.shields.io/github/license/hargeek/gin-auto-redoc)](./LICENSE)

## Installation

```bash
go get github.com/hargeek/gin-auto-redoc
```

## Usage

Import the package in your main file

```go
import _ "github.com/hargeek/gin-auto-redoc"
```

Register the Redoc middleware in your Gin application

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

## Configuration

You can set custom configuration by calling SetConfig before registering the middleware

```go
ginautodoc.SetConfig(ginautodoc.Config{
    SwaggerPath: "/custom/swagger",
})
ginautodoc.Register(r)
```