package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

type Student struct {
	Name string
	Age  int
}

func test(c *gin.Context) {

	c.HTML(http.StatusOK, "test.html", gin.H{
		"stuList": []Student{{Name: "yy", Age: 13}, {Name: "sdd", Age: 44}},
	})

}

func main() {
	router := gin.Default()

	router.SetFuncMap(template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
	}) // 自定义模板函数

	router.LoadHTMLGlob("template/*") // 加载模板

	router.GET("/test", test)

	router.Run() // 127.0.0.1:8080

}
