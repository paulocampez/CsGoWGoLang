package controllers

import (
	"fmt"
	"net/http"

	"../model"
	"github.com/gin-gonic/gin"
)

//Busca todos os items
func (db *Db) GetItems(c *gin.Context) {
	var items = new([]model.Item)

	var (
		result gin.H
	)
	//id := c.Param("id")
	db.DB.LogMode(true)
	db.DB.Find(&items)
	//db.DB.Where("id = ?", id).First(&item)
	result = gin.H{
		"result": items,
	}
	fmt.Println(result)
	c.JSON(http.StatusOK, result)
}

//Edita Item
func (db *Db) UpdatePerson(c *gin.Context) {
	var item model.Item
	id := c.Params.ByName("id")
	if err := db.DB.Where("id = ?", id).First(&item).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&item)
	db.DB.Save(&item)
	c.JSON(200, item)
}

func (db *Db) DeleteItem(c *gin.Context) {
	var item model.Item
	id := c.Params.ByName("id")
	db.DB.First(&item, id)
	db.DB.Delete(&item)
	c.JSON(http.StatusOK, gin.H{"success": "Item #" + id + " deleted"})
}

//Busca item por Id no banco
func (db *Db) GetItemsById(c *gin.Context) {
	var items = new([]model.Item)

	var (
		result gin.H
	)
	id := c.Param("id")
	db.DB.LogMode(true)
	db.DB.Where("id = ?", id).First(&items)
	result = gin.H{
		"result": items,
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
