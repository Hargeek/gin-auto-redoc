package ginautodoc

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {
	swaggerJSONPath := path.Join(globalConfig.SwaggerPath, "doc.json")
	redocPath := strings.Replace(globalConfig.SwaggerPath, "swagger", "doc", 1)
	registerRedoc(engine, swaggerJSONPath, redocPath)
}

func registerRedoc(engine *gin.Engine, swaggerJSONPath, redocPath string) {
	engine.GET(redocPath+"/*any", func(c *gin.Context) {
		handleRedoc(c, swaggerJSONPath)
	})
}

func handleRedoc(c *gin.Context, swaggerJSONPath string) {
	i := strings.LastIndex(c.Request.URL.Path, "/")
	if i == -1 {
		return
	}

	suffix := c.Request.URL.Path[i+1:]
	if suffix == "index.html" {
		redocHTML := fmt.Sprintf(`<!DOCTYPE html>
        <html>
          <head>
            <title>API Documentation</title>
            <meta charset="utf-8"/>
            <meta name="viewport" content="width=device-width, initial-scale=1">
            <link href="https://fonts.googleapis.com/css?family=Montserrat:300,400,700|Roboto:300,400,700" rel="stylesheet">
            <style>
              body {
                margin: 0;
                padding: 0;
              }
            </style>
          </head>
          <body>
            <redoc spec-url='%s'></redoc>
            <script src="https://cdn.redoc.ly/redoc/latest/bundles/redoc.standalone.js"></script>
          </body>
        </html>`, swaggerJSONPath)

		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, redocHTML)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"msg": "page not found",
	})
}
