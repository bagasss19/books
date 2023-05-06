package model

import "time"

type Book struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	ISBN        string    `json:"isbn"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
}

const TableBook = "books"

func (Book) TableBook() string {
	return TableBook
}
