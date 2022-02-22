package main

import (
	middleware "ginGO/MiddleWares"
	"ginGO/Routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.UserMidder{}.MiddleWare())

	// r.GET("/", func(c *gin.Context) {
	// 	c.String(200, "HEllo")
	// })
	Routers.Route(r)
	r.Run(":8022")

}
