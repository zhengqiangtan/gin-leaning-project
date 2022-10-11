package Test

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Age      int    `form:"age" json:"age" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	req := &User{}
	if err := c.Bind(req); err != nil {
		log.Printf("err:%v", err)
		c.JSON(http.StatusOK, gin.H{
			"errno":  "300",
			"errmsg": err.Error(),
		})
		return
	}

	if req.UserName != "admin" || req.Password != "123456" {
		c.JSON(http.StatusOK, gin.H{
			"errno":  "500",
			"errmsg": "password or username is wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errno":  "0",
		"errmsg": "login success",
	})
}
