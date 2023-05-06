package repository

import "github.com/books/service/books/model"

type GetAllBookOutput struct {
	Data []*model.Book

	// pagination
	Count   int64
	MaxPage int64
}

type PropsBooksGetAll struct {
	Title  string
	Author string
	ISBN   string

	// For Paggination
	Page  int64
	Limit int64

	// For Order
	Order string
	Sort  string
}
