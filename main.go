package main

import (
	"github.com/gin-gonic/gin"
)

type TestType struct {
	Name string
}

func TestTypeFunc(c *gin.Context) {
	h := TestType{Name: "Zuhair"}
	c.JSON(200, h)

}
func main() {
	//fmt.Println("Hello, World!")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"H": "Hello, World!",
		})
	})
	r.GET("/", TestTypeFunc)
	r.Run()
}
