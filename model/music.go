package model

type Music struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Title    string `json:"title" gorm:"not null"`
	Artist   string `json:"artist" gorm:"not null"`
	Album    string `json:"album"`
	Genre    string `json:"genre"`
	Duration int    `json:"duration"`
}
