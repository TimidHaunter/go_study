package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	// https://learnku.com/docs/gin-gonic/1.7/examples-rendering/11371
	r.GET("/string", returnString)
	r.GET("/string/notFound", returnStringNotFound)
	r.GET("/string/internalServerError", returnStringInternalServerError)

	r.GET("/json", returnJson)
	r.GET("/json/struct", returnJsonStruct)
	r.GET("/json/map", returnJsonMap)

	// https://learnku.com/docs/gin-gonic/1.7/examples-querystring-param/11388
	r.GET("/query/string", queryString)

	// https://learnku.com/docs/gin-gonic/1.7/examples-query-and-post-form/11369
	r.POST("/query/form", queryForm)

	// https://learnku.com/docs/gin-gonic/1.7/parameters-in-path/11407
	r.GET("/query/path/:id/:name", queryPath)

	r.Run()
}

// 返回一个字符串
func returnString(c *gin.Context) {
	c.String(200, "我是一个字符串")
}

// 返回一个http状态码是404的字符串
func returnStringNotFound(c *gin.Context) {
	c.String(404, "我是一个字符串,http状态码是404")
}

// 返回一个http状态码是500的字符串
func returnStringInternalServerError(c *gin.Context) {
	c.String(500, "我是一个字符串,http状态码是500")
}

// 返回一个JSON格式
func returnJson(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "我是一个JSON格式",
	})
}

// 使用结构体返回一个JSON格式
func returnJsonStruct(c *gin.Context) {
	// 定义一个返回结构体
	var msg struct {
		Name    string `json:"user"`
		Message string
		Age     int
	}
	// 实例化一个结构体
	msg.Name = "吕布"
	msg.Message = "虎牢关"
	msg.Age = 34

	c.JSON(http.StatusOK, msg)
}

// 使用Map返回一个JSON格式
func returnJsonMap(c *gin.Context) {
	// 定义一个map
	data := map[string]interface{}{
		"name":    "董卓",
		"message": "戏貂蝉",
		"age":     56,
	}
	c.JSON(http.StatusOK, data)
}

// 查询字符串参数
func queryString(c *gin.Context) {
	id := c.Query("id")
	user := c.DefaultQuery("user", "Guest")
	address := c.DefaultQuery("address", "地球")

	var msg struct {
		Id      string
		User    string
		Address string
	}

	msg.Id = id
	msg.User = user
	msg.Address = address

	c.JSON(http.StatusOK, msg)
}

// Form表单
func queryForm(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")
	message := c.PostForm("message")

	var msg struct {
		Id      string
		Page    string
		Name    string
		Message string
	}

	msg.Id = id
	msg.Page = page
	msg.Name = name
	msg.Message = message

	c.JSON(http.StatusOK, msg)
}

// 获取路径中的参数
func queryPath(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")

	var msg struct {
		Id   string `json:"id"`
		Name string `json:"user"`
	}

	msg.Id = id
	msg.Name = name

	c.JSON(http.StatusOK, msg)
}
