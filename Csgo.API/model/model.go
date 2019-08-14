package model

import "github.com/jinzhu/gorm"

//Model Item
type Item struct {
	gorm.Model
	Name string
}

//Model TransformedItem
type TransformedItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
