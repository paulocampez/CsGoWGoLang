package main

import (
	//model "../Csgo.API/model"
	service "../Csgo.API/services"
	//"fmt"
	//"github.com/gin-gonic/gin"
	//"net/http"
)

func main() {
	//fmt.Println(service.CheckDeaths(service.GetByRound(1)))
	//fmt.Println(service.GetAllKillsInRound(21))
	//fmt.Println(service.GetPlayersByRound(2))
	service.GetParser()
	//fmt.Println(service.CheckDeathsByRound(service.GetByRound(1)))
	//checkGamesQt()

	// db := config.DBInit()
	// router := gin.Default()
	// Db := &controllers.Db{DB: db}

	// router.GET("/items/:id", Db.GetItemsById)
	// router.GET("/items", Db.GetItems)
	// router.POST("/items", Db.CreateItem)
	// router.PUT("/items", Db.UpdatePerson)
	// router.DELETE("/items/:id", Db.DeleteItem)
	// router.Run(":8080")
	// fmt.Println(db)

}
