package Routers

import (
	"ginGO/AdminController"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Route(r *gin.Engine) {
	Admin := r.Group("/admin")
	{
		Admin.GET("/", AdminController.Controllers{}.AdminHome)
		Admin.GET("/page/:name", AdminController.Controllers{}.UpdatePage)
		Admin.GET("/admin/:age", func(c *gin.Context) {
			// var b strings.Builder
			val := c.Param("age")
			_, err := strconv.Atoi(val)
			if err != nil {
				c.JSON(200, gin.H{"status": err})
			}
		})
	}
}
