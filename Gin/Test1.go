package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

/**
gin框架的请求测试
*/
func main() {
	r := gin.Default() // 生成了一个实例，这个实例即 WSGI 应用程序

	// case1: 无参数
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Gin")
	})

	// case2:动态的路由
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	//r.GET("/user/:age", func(c *gin.Context) {
	//	name := c.Param("age")
	//	c.String(http.StatusOK, "age is %s", name)
	//})

	// case3:获取query参数，匹配users?name=xxx&role=xxx，role可选
	r.GET("/users", func(c *gin.Context) {
		name := c.Query("name")
		role := c.DefaultQuery("role", "teacher")
		c.String(http.StatusOK, "%s is a %s", name, role)
	})

	// case4: POST示例
	// linux: curl http://localhost:9999/form  -X POST -d 'username=tzq&password=1234'
	// win:   curl -Uri 'http://localhost:9999/form' -Body 'username=tzq&password=1234' -Method 'POST'
	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "123456") // 可设置默认值

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})
	//  {"password":"1234","username":"tzq"}

	// case5: GET 和 POST 混合
	r.POST("/posts", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		username := c.PostForm("username")
		password := c.DefaultPostForm("username", "123456") // 可设置默认值

		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"page":     page,
			"username": username,
			"password": password,
		})
	})
	// {"id":"100","page":"10","password":"tzq","username":"tzq"}

	//case6: map参数
	// $ curl -g "http://localhost:9999/post?ids[Jack]=001&ids[Tom]=002" -X POST -d 'names[a]=Sam&names[b]=David'
	r.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})

	// case7: 分组路由 （重要）
	// group routes 分组路由
	defaultHandler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"path": c.FullPath(),
		})
	}
	// group: v1
	v1 := r.Group("/v1")
	{
		v1.GET("/posts", defaultHandler)
		v1.GET("/series", defaultHandler)
	}
	// group: v2
	v2 := r.Group("/v2")
	{
		v2.GET("/posts", defaultHandler)
		v2.GET("/series", defaultHandler)
	}

	//自定义日志文件:写入文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	gin.SetMode("debug")
	gin.DisableConsoleColor()
	r.Run(":9999")

}
