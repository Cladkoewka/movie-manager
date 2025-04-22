package model

import "time"

type MoviePoster struct {
	ID        int64     `json:"id"`
	MovieID   int64     `json:"movie_id"`
	Poster    []byte    `json:"poster"`    
	MimeType  string    `json:"mime_type"` 
	CreatedAt time.Time `json:"created_at"`
}
