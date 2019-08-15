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
	TotalKill int
	Players   []Player
	Kills     []Kills
}

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
