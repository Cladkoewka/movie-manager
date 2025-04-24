package model

type Review struct {
	ID int64 `json:"id"`
	MovieID int64 `json:"movie_id"` 
	Comment string `json:"comment"`
}