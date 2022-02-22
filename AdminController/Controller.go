package AdminController

import (
	"github.com/gin-gonic/gin"
)

type Controllers struct {
}

func (con Controllers) AdminHome(c *gin.Context) {
	c.String(200, "我是首页")
}

func (con Controllers) UpdatePage(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, gin.H{
		"结果": name,
	})

}
