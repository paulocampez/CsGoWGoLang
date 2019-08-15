package services

import (
	model "../model"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func PlayerGetScore() {

}

func TotalKill(kill int) int {
	kill++
	return kill
}

func ReadFile() []string {
	//usar esse comando quando estiver no servidor aws
	dat, _ := ioutil.ReadFile("games.log")
	//dat, _ := ioutil.ReadFile(basepath + "\\logs\\games.log")
	trimSpaces := strings.TrimSpace(string(dat))
	lstLines := strings.Split(trimSpaces, "\n")
	//fmt.Println(lstLines)
	return lstLines
}

func CheckDeaths(lines []string) model.Deaths {
	var deathsMajor model.Deaths
	for _, element := range lines {
		strRegex := `[ˆ0-9]: (.*) killed (.*) [ˆb,y]`
		match, _ := regexp.MatchString(strRegex, element)
		if match {
			r, _ := regexp.Compile(strRegex)
			nickname := r.FindStringSubmatch(element)
			deathsMajor.DeathsFather = append(deathsMajor.DeathsFather, model.DeathGame{Player1: string(nickname[1]), Player2: string(nickname[2])})
		}
	}
	return deathsMajor
}

func AddPlayer(line string) string {
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

func RemoveDuplicates(elements []string) []string {
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

func GetPlayersByRound(round int) []string {
	lines := GetByRound(round)
	var player []string
	for _, element := range lines {
		if strings.Contains(element, "ClientUserinfoChanged") {
			player = append(player, AddPlayer(element))
		}
	}
	player = RemoveDuplicates(player)
	return player
}

func GetByRound(round int) []string {
	lstLines := ReadFile()
	round = round - 1
	var rounds []string
	line := Find(lstLines, "InitGame")
	if round == len(line)-1 {
		//fmt.Println("ESTOU NO ULTIMO ROUND")
		//fmt.Println(lstLines[line[round-1]:len(lstLines)])
		rounds = lstLines[line[round-1]:len(lstLines)]
	} else {
		rounds = lstLines[line[round]:line[round+1]]
	}
	//fmt.Println(rounds)
	return rounds
}

func CheckDeathsByRound(lines []string) int {
	//lines = GetByRound(21)
	chkdeaths := CheckDeaths(lines)
	return len(chkdeaths.DeathsFather)
}

func GetParser() {
	numberRounds := checkGamesQt()
	fmt.Println("Numero de Rounds: ", numberRounds)
	//var allGames model.Root
	for i := 1; i <= numberRounds; i++ {
		fmt.Println("--------------------------------")
		fmt.Println("Players in Game", i, ":")
		justSplit := strings.Join(GetPlayersByRound(i), ", ")
		fmt.Println(justSplit)
		fmt.Println("Total Kills in Game", CheckDeathsByRound(GetByRound(i)))

		for _, element := range GetPlayersByRound(i) {
			fmt.Println("Player", element, "Score", GetKillByPlayerAndRound(element, i))
		}
		fmt.Println("--------------------------------")
	}
	//	fmt.Println(allGames.Games)
}

func GetKillByPlayerAndRound(player string, round int) int {
	deathsMajor := CheckDeaths(GetByRound(round))
	var score int
	for _, element := range deathsMajor.DeathsFather {
		if element.Player1 == player {
			score++
		}
	}
	//fmt.Println("Player", player, "Score:", score)
	return score
}

func checkGamesQt() int {
	lines := ReadFile()
	qtt := 0
	for _, element := range lines {
		if strings.Contains(element, "InitGame") {
			qtt++
		}
	}
	return qtt
}

func PlayerScore(players []string) bool {
	return true
}
