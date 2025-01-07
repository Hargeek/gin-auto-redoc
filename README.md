# gin-auto-redoc

A Go package to automatically register Redoc documentation for Gin applications.

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