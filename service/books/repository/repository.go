package repository

import (
	"context"

	"github.com/books/service/books/model"
	"github.com/books/utils"
	"gorm.io/gorm"
)

type BookRepository interface {
	GetAllBook(ctx context.Context, props PropsBooksGetAll) (resp GetAllBookOutput, err error)
	GetOneBook(ctx context.Context, ID string) (data *model.Book, err error)
	UpdateBook(ctx context.Context, props model.Book, fields []string, ID string) (err error)
	CreateBook(ctx context.Context, props model.Book) (ID *string, err error)
	DeleteBook(ctx context.Context, ID string) error
}

type bookRepository struct {
	repoPrefix string
	db         *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{
		repoPrefix: utils.RepositoryPrefix + "booksrepository",
		db:         db,
	}
}
