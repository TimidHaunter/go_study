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
