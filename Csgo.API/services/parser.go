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

func GetPlayersByRound(round int) []model.Player {
	lines := GetByRound(round)
	var player []model.Player
	var playerTransformed []model.Player
	for _, element := range lines {
		if strings.Contains(element, "ClientUserinfoChanged") {
			player = append(player, model.Player{Name: AddPlayer(element)})
		}
	}
	var s []string
	for _, v := range player {
		s = append(s, v.Name)
	}

	s = RemoveDuplicates(s)

	for _, t := range s {
		playerTransformed = append(playerTransformed, model.Player{Name: t})
	}

	return playerTransformed
}

func GetByRound(round int) []string {
	lstLines := ReadFile()
	round = round - 1
	var rounds []string
	line := Find(lstLines, "InitGame")
	if round == len(line)-1 {
		//fmt.Println(lstLines[line[round-1]:len(lstLines)])
		rounds = lstLines[line[round-1]:len(lstLines)]
	} else {
		rounds = lstLines[line[round]:line[round+1]]
	}
	return rounds
}

func CheckDeathsByRound(lines []string) int {
	chkdeaths := CheckDeaths(lines)
	return len(chkdeaths.DeathsFather)
}

func GetParser() model.Root {
	numberRounds := checkGamesQt()

	var allGames model.Root
	fmt.Println("Numero de Rounds: ", numberRounds)
	for i := 1; i <= numberRounds; i++ {

		fmt.Println("Jogo : ", i)
		fmt.Println("Numero de kills no jogo: ", CheckDeathsByRound(GetByRound(i)))
		fmt.Println("Jogadores : ", GetPlayersByRound(i))
		fmt.Println("Score dos Jogadores: ", GetAllKillsInRound(i))

		allGames.Games = append(allGames.Games, model.Game{GameNumber: i, TotalKill: CheckDeathsByRound(GetByRound(i)), Players: GetPlayersByRound(i), Kills: GetAllKillsInRound(i)})
	}
	return allGames
}

func GetAllKillsInRound(round int) []model.Kills {
	var lstKills []model.Kills
	for _, element := range GetPlayersByRound(round) {
		mdl := GetKillByPlayerAndRound(element.Name, round)
		lstKills = append(lstKills, model.Kills{Player: mdl.Player, Score: mdl.Score})
	}
	return lstKills
}

func GetKillByPlayerAndRound(player string, round int) model.Kills {
	deathsMajor := CheckDeaths(GetByRound(round))
	var playerScored model.Kills
	var score int
	for _, element := range deathsMajor.DeathsFather {
		if element.Player1 == player {
			score++
		}
	}
	playerScored = model.Kills{Player: player, Score: score}
	return playerScored
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
