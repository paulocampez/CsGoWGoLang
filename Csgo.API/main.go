package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Deaths struct {
	deathsFather []DeathGame
}

type DeathGame struct {
	player1 string
	player2 string
}

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func readFile() []string {
	dat, _ := ioutil.ReadFile(basepath + "\\logs\\games.log")

	trimSpaces := strings.TrimSpace(string(dat))
	lstLines := strings.Split(trimSpaces, "\n")
	fmt.Println(strings.Contains(lstLines[1], "InitGame"))
	return lstLines
}

func checkDeaths(lines []string) Deaths {
	var deathsMajor Deaths
	for _, element := range lines {
		strRegex := "[ˆ0-9]: (.*) killed (.*) [ˆb,y]"
		match, _ := regexp.MatchString(strRegex, element)
		if match {

			r, _ := regexp.Compile(strRegex)
			nickname := r.FindStringSubmatch(element)
			deathsMajor.deathsFather = append(deathsMajor.deathsFather, DeathGame{player1: string(nickname[1]), player2: string(nickname[2])})
		}
	}
	return deathsMajor
}

func checkGamesQt() int {
	lines := readFile()
	qtt := 0
	for _, element := range lines {
		if strings.Contains(element, "InitGame") {
			qtt++
		}
	}
	fmt.Println(qtt)
	return qtt
}

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {

	checkGamesQt()

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

	//Login usando jwt
	//router.POST("/login", LoginHandler)
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
