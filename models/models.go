package models

import (
	"database/sql"
	"time"
)

// wrapper for database
type Models struct {
	DB DBModel
}

// returns models with db pool
func NewModels(db *sql.DB) Models {
	return Models {
		DB: DBModel{DB: db},
	}
}

// type for Meme
type Meme struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Rating int `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// type for User
type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
	Memes []Meme `json:"memes"`
	UserId int `json:user_id`
}

// type for scraping result
type ScrapingResult struct {
	MemeTitle string `json:"meme_title"`
	Image string `json:"image"`
	About string `json:"about"`
	AboutText string `json:"about_text"`
	Origin string `json:"origin"`
	OriginText string `json:"origin_text"`
	Spread string `json:"spread"`
	SpreadText string `json:"spread_text"`
}