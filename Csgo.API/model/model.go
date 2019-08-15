package model

import "github.com/jinzhu/gorm"

//Model Item
type Item struct {
	gorm.Model
	Name string
}

type TransformedItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Deaths struct {
	DeathsFather []DeathGame
}

type DeathGame struct {
	Player1 string
	Player2 string
}

type Player struct {
	Name string
}

type PlayerScore struct {
	Name  string
	Score int
}

type Kills struct {
	Player string
	Score  int
}
type Game struct {
	GameNumber int      `json:"game"`
	TotalKill  int      `json:total_kills`
	Players    []Player `json:players`
	Kills      []Kills  `json:kills`
}
type Root struct {
	Games []Game
}

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//game_1: {
//     total_kills: 45;
//     players: ["Dono da bola", "Isgalamido", "Zeh"]
//  kills: {
//       "Dono da bola": 5,
//       "Isgalamido": 18,
//       "Zeh": 20
//   }
//}
