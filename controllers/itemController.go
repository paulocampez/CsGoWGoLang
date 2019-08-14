package controllers

import (
	"fmt"
	"net/http"

	"../model"
	"github.com/gin-gonic/gin"
)

//Busca item por Id no banco
func (db *Db) GetItem(c *gin.Context) {
	var items = new([]model.Item)
	var _items = new([]model.TransformedItem)

	var (
		//	item   model.Item
		result gin.H
	)
	id := c.Param("id")
	fmt.Println("O ID PASSADO FOI O NUMERO :", id)
	db.DB.LogMode(true)
	db.DB.Find(&items)

	//db.DB.Where("id = ?", id).First(&item)
	fmt.Println("PASSOU NO FIND")

	_items = append(_items, model.TransformedItem{ID: items.id, Name: items.name})

	fmt.Println(_items)

	result = gin.H{
		"result": _items,
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
