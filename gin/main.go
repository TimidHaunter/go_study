package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	r.GET("/string", returnString)
	r.GET("/string/notFound", returnStringNotFound)
	r.GET("/string/internalServerError", returnStringInternalServerError)

	r.GET("/json", returnJson)

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
