English | [简体中文](./README-cn.md) 

# gin-auto-redoc

A Go package to automatically register Redoc documentation for a Gin application that already includes Swagger documentation

![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.22-%23007d9c)
![Gin Version](https://img.shields.io/badge/Gin-%3E%3D1.10-green)
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
    ginautodoc "github.com/hargeek/gin-auto-redoc"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.New()
    ginautodoc.Register(r)
    r.Run()
}
```

## Configuration

### Customize Swagger Route

By default, the Swagger Documentation is available at /api/v1/swagger and Redoc documentation is available at /api/v1/doc. To customize these paths, call SetConfig with your desired configuration before registering the middleware.

```go
ginautodoc.SetConfig(ginautodoc.Config{
    SwaggerPath: "/custom/swagger",
})
ginautodoc.Register(r)
```

### Reverse Proxy

If your documentation is served through a reverse proxy and the path is rewritten, you need to pass the X-Forwarded-Prefix header to the service, otherwise you will not be able to access it, for example

- nginx configuration example:

```nginx
location /custom_prefix/ {```
    proxy_pass http://localhost:8080;
    proxy_set_header X-Forwarded-Prefix "/custom_prefix";
}
```

- Kubernetes ingress configuration example:

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: my-ingress
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /$1
    nginx.ingress.kubernetes.io/use-regex: "true"
    nginx.ingress.kubernetes.io/configuration-snippet: |
      proxy_set_header X-Forwarded-Proto $http_x_forwarded_proto;
      proxy_set_header X-Forwarded-Prefix "/custom_prefix";
spec:
  ingressClassName: nginx
  rules:
  - host: example.com
    http:
      paths:
      - backend:
          service:
            name: api-service
            port:
              number: 8080
        path: /custom_prefix/(.*)$
        pathType: Prefix
```

## Accessing Redoc

Once you have started your Gin application, you can access the Redoc documentation at `http://localhost:<port>/api/v1/doc/index.html`
