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

type Player struct {
	name string
}
type Kills struct {
	player string
	score  int
}
type Game struct {
	totalKill int
	players   []Player
	kills     []Kills
}

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func playerGetScore() {

}

func totalKill(kill int) int {
	kill++
	return kill
}

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

func addPlayer(line string) string {
	strRegex := `n\\(.*)\\t\\`
	r, _ := regexp.Compile(strRegex)
	match, _ := regexp.MatchString(strRegex, line)
	if match {
		nickname := r.FindStringSubmatch(line)
		return nickname[1]
	}
	return ""
}

//metodo para buscar string em lista
func Find(a []string, x string) []int {
	var lstArray []int
	for i, n := range a {
		if strings.Contains(n, x) {
			lstArray = append(lstArray, i)
		}
	}
	return lstArray
}

func removeDuplicates(elements []string) []string {
	encountered := map[string]bool{}

	for v := range elements {
		encountered[elements[v]] = true
	}

	result := []string{}
	for key, _ := range encountered {
		result = append(result, key)
	}
	return result
}

func getPlayersByRound(round int) []string {
	lines := getByRound(round)
	var player []string
	for _, element := range lines {
		if strings.Contains(element, "ClientUserinfoChanged") {
			player = append(player, addPlayer(element))
		}
	}
	player = removeDuplicates(player)
	fmt.Println(player)
	return player
}

func getByRound(round int) []string {
	lstLines := readFile()
	var rounds []string
	line := Find(lstLines, "InitGame")
	if round == len(line) {
		fmt.Println("ESTOU NO ULTIMO ROUND")
		//fmt.Println(lstLines[line[round-1]:len(lstLines)])
		rounds = lstLines[line[round-1]:len(lstLines)]
	} else {
		rounds = lstLines[line[round]:line[round+1]]
	}
	//fmt.Println(rounds)
	return rounds
}

func getParser() {
	var game Game
	totalGames := checkGamesQt()
	for i := 1; i < totalGames; i++ {

	}
	fmt.Println(game)
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
	getPlayersByRound(1)
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
