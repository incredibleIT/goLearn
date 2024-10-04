package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Article struct {
	Content string
	Author  string
}

func f1(c *gin.Context) {
	c.JSON(200, gin.H{
		"请求方式": c.Request.Method,
	})
}

func f2(c *gin.Context) {

	fmt.Println("GET:")
	fmt.Println(c.Query("year"), c.Query("month"))
	fmt.Println("POST:")

	// urlencoded格式GIN会主动处理, 用PostForm()可取到
	fmt.Println(c.PostForm("year"), c.PostForm("month"))

	// 对于json格式GIN不会处理,我么需要自己处理, 由c.Request.Body拿到原始的请求体读取并解析
	var resData map[string]int
	//data, _ := ioutil.ReadAll(c.Request.Body)
	//json.Unmarshal(data, &resData)
	//fmt.Println(resData["year"])
	//fmt.Println(resData["month"])

	c.BindJSON(&resData) // 直接解析JSON
	fmt.Println(resData)

	c.JSON(200, gin.H{
		"mess": "hello",
	})
}

func noRoute(c *gin.Context) {
	c.JSON(200, gin.H{
		"err": "没有匹配",
	})
}

func article(c *gin.Context) {
	// 获取路径参数
	fmt.Println(c.Param("year"), c.Param("month"))

	// 数据库中查询文章
	article1 := Article{"1", "1"}
	article2 := Article{"2", "2"}
	article3 := Article{"3", "3"}
	articleList := []Article{article1, article2, article3}
	fmt.Println(articleList)

	// 重定向用: c.Redirect(302, "")

	c.JSON(200, gin.H{
		"articleList": articleList,
	})
}

func main() {

	router := gin.Default()

	// 不管什么类型的请求都走
	//router.Any("/f1", f1)

	// 其他路由没有匹配成功的走noROute
	router.NoRoute(f1)

	// 路由组
	//gr := router.Group("/v1")
	// 然后gr为此组的路由由他进行分发

	// 获取路由参数(1. 请求参数, 2. 路径参数)
	// get请求参数
	router.GET("/book", f2)

	// post 请求参数, post请求体参数主流分为两种格式urlencoded 和 json
	router.POST("/book", f2)

	// 路径参数
	router.GET("/article/:year/:month", article)

	router.Run(":8080")

}
