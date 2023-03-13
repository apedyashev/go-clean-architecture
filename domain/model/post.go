package model

type Post struct {
	ID     uint     `json:"id"`
	Title  string   `json:"title"`
	Images []string `json:"images"`
}

func (Post) TableName() string { return "posts" }
