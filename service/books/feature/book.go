package feature

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/books/service/books/model"
	"github.com/books/service/books/repository"
	"github.com/books/utils"
	"github.com/google/uuid"
	"github.com/spf13/cast"
)

func (b *bookFeature) GetOneBook(ctx context.Context, ID string) (data *model.Book, err error) {
	prefix := b.featurePrefix + ".GetOneBook"

	data, err = b.bookRepo.GetOneBook(ctx, ID)
	if err != nil {
		if err.Error() == utils.ErrorDataNotFound {
			return nil, errors.New(utils.ErrorDataNotFound)
		}
		return nil, fmt.Errorf("[%s] error while get one book: %+v", prefix, err)
	}

	return data, nil
}

func (b *bookFeature) GetAllBook(ctx context.Context, props PropsBooksGetAll) (resp GetAllBookResponse, err error) {
	prefix := b.featurePrefix + ".GetAllBook"

	propsGetAll := repository.PropsBooksGetAll{
		Title:  props.Title,
		Author: props.Author,
		ISBN:   props.ISBN,
		Page:   props.Page,
		Limit:  props.Limit,
		Order:  props.Order,
		Sort:   props.Sort,
	}
	output, err := b.bookRepo.GetAllBook(ctx, propsGetAll)
	if err != nil {
		err = fmt.Errorf("[%s] error while get all book: %+v", prefix, err)
		return
	}
	resp.Books = output.Data
	resp.MaxPage = output.MaxPage
	resp.Count = output.Count

	return
}

func (b *bookFeature) CreateBook(ctx context.Context, props BookPayload) (ID *string, err error) {
	prefix := b.featurePrefix + ".CreateBook"

	propsCreate := model.Book{
		ID:          uuid.NewString(),
		Title:       props.Title,
		ISBN:        props.ISBN,
		Author:      props.Author,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}
	ID, err = b.bookRepo.CreateBook(ctx, propsCreate)
	if err != nil {
		return nil, fmt.Errorf("[%s] error while create book: %+v", prefix, err)
	}

	return ID, nil
}

func (b *bookFeature) UpdateBookByFields(ctx context.Context, props PropsBookUpdateByFields, ID string) error {
	prefix := b.featurePrefix + ".UpdateBookByFields"
	var (
		fields      []string
		updatedBook model.Book
	)

	for field, val := range props.Data {
		switch field {
		case "title":
			newValue, err := cast.ToStringE(val)
			if err != nil {
				continue
			}
			updatedBook.Title = newValue
			fields = append(fields, field)

		case "author":
			newValue, err := cast.ToStringE(val)
			if err != nil {
				continue
			}
			updatedBook.Author = newValue
			fields = append(fields, field)

		case "isbn":
			newValue, err := cast.ToStringE(val)
			if err != nil {
				continue
			}
			updatedBook.ISBN = newValue
			fields = append(fields, field)

		}
	}

	if len(fields) > 0 {
		updatedBook.UpdatedTime = time.Now()
		fields = append(fields, "updated_time")
		if err := b.bookRepo.UpdateBook(ctx, updatedBook, fields, ID); err != nil {
			if err.Error() == utils.ErrorDataNotFound {
				return errors.New(utils.ErrorDataNotFound)
			}
			return fmt.Errorf("[%s] error while UpdateBookByFields book: %+v", prefix, err)
		}
	}

	return nil
}

func (b *bookFeature) DeleteBook(ctx context.Context, ID string) error {
	prefix := b.featurePrefix + ".DeleteBook"

	// Check if data is exist
	_, err := b.bookRepo.GetOneBook(ctx, ID)
	if err != nil {
		if err.Error() == utils.ErrorDataNotFound {
			return errors.New(utils.ErrorDataNotFound)
		}
		return fmt.Errorf("[%s] error while get one book: %+v", prefix, err)
	}

	err = b.bookRepo.DeleteBook(ctx, ID)
	if err != nil {
		return fmt.Errorf("[%s] error while delete book: %+v", prefix, err)
	}

	return nil
}
