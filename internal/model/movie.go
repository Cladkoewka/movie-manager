package model

import "time"

type Movie struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	ReleaseDate time.Time `json:"release_date"`
	Genre string `json:"genre"`
	Director string `json:"director"`
	Rating float64 `json:"rating"`
	Duration int `json:"duration"`
	Language string `json:"language"`
	PosterURL string `json:"poster_url"`
	TrailerURL string `json:"trailer_url"`
}