package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义接收数据的结构体
type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User     string `uri:"user" binding:"required"`
	Password string `uri:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	// Uri绑定
	r.GET("/:user/:password", func(c *gin.Context) {
		var login Login
		// 将request中的param解析到结构体
		if err := c.ShouldBindUri(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if login.User != "root" || login.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "incorrect user or password"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "login success"})
	})
	r.Run(":8001")
}
