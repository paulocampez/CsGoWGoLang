package main

import (
	"fmt"
	"net/http"

	"./config"
	"./controllers"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {

	//configuracoes de conexao
	db := config.DBInit()
	router := gin.Default()
	Db := &controllers.Db{DB: db}


	//Login usando jwt
	//router.POST("/login", LoginHandler)

	//Metodo de create
	//TODO: Não está passando o item correto usando Postman
	router.POST("/items", Db.CreateItem)
	
	//Metodo get por Id
	router.GET("/items/:id", Db.GetItem)
	router.Run(":3000")
	//adicionado println para tirar erro da variavel
	fmt.Println(db)
}

//Codigo retirado do https://godoc.org/
func LoginHandler(c *gin.Context) {
	var user Credential
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "can't bind struct",
		})
	}
	if user.Username != "myname" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "wrong username or password",
		})
	} else {
		if user.Password != "myname123" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "wrong username or password",
			})
		}
	}
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

//Codigo retirado do https://godoc.org/
func Auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	// if token.Valid && err == nil {
	if token != nil && err == nil {
		fmt.Println("token verified")
	} else {
		result := gin.H{
			"message": "not authorized",
			"error":   err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}
