package feature

import (
	"context"

	"github.com/books/service/books/model"
	"github.com/books/service/books/repository"
	"github.com/books/utils"
)

type BookFeature interface {
	GetOneBook(ctx context.Context, ID string) (data *model.Book, err error)
	GetAllBook(ctx context.Context, props PropsBooksGetAll) (resp GetAllBookResponse, err error)
	CreateBook(ctx context.Context, props BookPayload) (ID *string, err error)
	UpdateBookByFields(ctx context.Context, props PropsBookUpdateByFields, ID string) error
	DeleteBook(ctx context.Context, ID string) error
}

type bookFeature struct {
	featurePrefix string
	bookRepo      repository.BookRepository
}

func NewBookFeature(bookRepo repository.BookRepository) BookFeature {
	return &bookFeature{
		featurePrefix: utils.RepositoryPrefix + "booksfeature",
		bookRepo:      bookRepo,
	}
}
