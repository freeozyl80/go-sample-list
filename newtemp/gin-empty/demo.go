package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.POST("/", home)
	r.Run()
}

func home(c *gin.Context) {
	type name struct {
		N *int `json:"n" form:"n" binding:"required"`
	}
	var n name
	err := c.ShouldBindJSON(&n)
	if err != nil {
		c.Writer.WriteString(err.Error())
		return
	}
	c.Writer.WriteString("ok")
}

// ok
// {
//	"n":0
// }

// fail
// {}
