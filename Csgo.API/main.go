package main

import (
	//model "../Csgo.API/model"
	service "../Csgo.API/services"
	//"fmt"
	"github.com/gin-gonic/gin"
	//"net/http"
)

func main() {
	var root = service.GetParser()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": root,
		})
	})
	r.Run(":8080")

}
