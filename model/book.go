package model

type Book struct {
	ID     string `gorm: autoIncrement; json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
