package model

type Comment struct {
	CommentsID int `json:"comments_id"`
	UseID uint64 `json:"use_id"`
	MovieID int `json:"movie_id"`
	Comments string
}
