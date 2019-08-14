package controllers

import (
	"fmt"
	"net/http"

	"../model"
	"github.com/gin-gonic/gin"
)

//Busca item por Id no banco
func (db *Db) GetItem(c *gin.Context) {
	var (
		item   model.Item
		result gin.H
	)
	id := c.Param("id")
	db.DB.LogMode(true)
	db.DB.Debug().Find(&item)

	db.DB.Where("id = ?", id).First(&item)
	fmt.Println(db.DB.Find(&item))

	result = gin.H{
		"result": item,
		"count":  1,
	}
	fmt.Println(result)
	c.JSON(http.StatusOK, result)
}
