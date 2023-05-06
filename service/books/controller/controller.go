package controller

import (
	"net/http"

	"github.com/books/service/books/feature"
	"github.com/books/utils"
)

type BookController interface {
	GetAllBook(w http.ResponseWriter, r *http.Request)
	GetOneBook(w http.ResponseWriter, r *http.Request)
	CreateBook(w http.ResponseWriter, r *http.Request)
	DeleteBook(w http.ResponseWriter, r *http.Request)
	UpdateBook(w http.ResponseWriter, r *http.Request)
}

type bookController struct {
	*utils.App
	bookFeature feature.BookFeature
}

func NewBookController(app *utils.App, bookFeature feature.BookFeature) BookController {
	return &bookController{
		App:         app,
		bookFeature: bookFeature,
	}
}
