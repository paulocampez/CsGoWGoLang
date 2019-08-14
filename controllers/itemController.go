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

func (db *Db) CreateItem(c *gin.Context) {

	//criar jsondecoder
	var (
		item   model.Item
		result gin.H
	)
	name := c.PostForm("name")
	name2 := c.PostForm("Name")
	fmt.Println("Nomes: ", name, name2)
	item.Name = name
	db.DB.LogMode(true)
	db.DB.Debug().Create(&item)
	db.DB.Create(&item)
	result = gin.H{
		"result": item,
	}
	c.JSON(http.StatusOK, result)
}
