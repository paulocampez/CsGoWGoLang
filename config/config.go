package config

import (
	"fmt"

	"../model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var (
	server   = "dbledark.database.windows.net"
	port     = 1433
	user     = "Ledark"
	password = "C4bymbxwo!"
	database = "Ledark"
)

// DBInit cria conexao com o banco
func DBInit() *gorm.DB {
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		server, user, password, port, database)
	db, err := gorm.Open("mssql", connectionString)
	if err != nil {
		fmt.Println(err)
		panic("falha ao conectar na base")
	}

	db.AutoMigrate(model.Item{})
	return db
}
