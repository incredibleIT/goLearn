package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func login(c *gin.Context) {

	// Gin上下文对象c 相当于w and r
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "login.html", nil)
	} else { // post
		username := c.PostForm("username")
		password := c.PostForm("pwd")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	}
}

func del(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"type": "del",
	})

}

func put(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"type": "put",
	})
}

func main() {

	// 定义Gin路由分发对象
	router := gin.Default()
	// 分发路由 符合restful规范
	router.GET("/login", login)
	router.POST("/login", login)
	router.DELETE("/del", del)
	router.PUT("/put", put)
	// 加载模板资源
	router.LoadHTMLGlob("./template/*")

	// 在指定端口监听并启动Serve
	_ = router.Run(":8080")

}
