package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	router := gin.New()
	// LoggerWithFormatter 中间件会写入日志到 gin.DefaultWriter
	// 默认 gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 你的自定义格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.Run(":8080")
}

//::1 - [Thu, 13 Oct 2022 17:25:13 CST] "GET /ping HTTP/1.1 200 0s "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36" "
