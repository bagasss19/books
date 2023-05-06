package repository

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/books/service/books/model"
	"github.com/books/utils"
	"gorm.io/gorm"
)

func (b *bookRepository) GetAllBook(ctx context.Context, props PropsBooksGetAll) (resp GetAllBookOutput, err error) {
	var (
		prefix     = b.repoPrefix + ".GetAllBook"
		query      = b.db.WithContext(ctx)
		queryCount = b.db.WithContext(ctx)
		data       []*model.Book
		total      int64
		wg         sync.WaitGroup
	)

	query = query.Order(fmt.Sprintf("%s %s", props.Order, props.Sort))

	// Selector
	if props.Page != 0 && props.Limit != 0 {
		query = utils.FilterPaginate(query, props.Page, props.Limit)
	}

	if props.Title != "" {
		query = filterBookByTitle(query, props.Title)
	}

	if props.ISBN != "" {
		query = filterBookByISBN(query, props.ISBN)
	}

	if props.Author != "" {
		query = filterBookByAuthor(query, props.Author)
	}

	wg.Add(2)

	go func() {
		defer wg.Done()
		// execute
		if err = query.Find(&data).Error; err != nil {
			err = fmt.Errorf("[%s] error while execute query: %+v", prefix, err)
			return
		}

		resp.Data = data

	}()

	go func() {
		defer wg.Done()
		// execute
		if err = queryCount.Find(&data).Count(&total).Error; err != nil {
			err = fmt.Errorf("[%s] error while execute query: %+v", prefix, err)
			return
		}

		resp.Count = total
		resp.MaxPage = utils.CountMaxPage(resp.Count, props.Limit)
	}()

	wg.Wait()

	return
}

func (b *bookRepository) GetOneBook(ctx context.Context, ID string) (data *model.Book, err error) {
	prefix := b.repoPrefix + ".GetOneBook"
	query := b.db.WithContext(ctx)

	if err := query.Where("id = ?", ID).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New(utils.ErrorDataNotFound)
		}
		return nil, fmt.Errorf("[%s] error while execute query: %+v", prefix, err)
	}

	return data, nil
}

func (b *bookRepository) UpdateBook(ctx context.Context, props model.Book, fields []string, ID string) (err error) {
	prefix := b.repoPrefix + ".UpdateBook"

	query := b.db.WithContext(ctx).Select(fields)
	query = filterBookByID(query, ID)

	if err := query.Updates(&props).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New(utils.ErrorDataNotFound)
		}
		return fmt.Errorf("[%s] error while execute query: %+v", prefix, err)
	}

	return nil
}

func (b *bookRepository) CreateBook(ctx context.Context, props model.Book) (ID *string, err error) {
	prefix := b.repoPrefix + ".CreateBook"
	query := b.db.WithContext(ctx)

	db := query.Create(&props)
	if err = db.Error; err != nil {
		return nil, fmt.Errorf("[%s] %+v", prefix, err)
	}
	return &props.ID, nil
}

func (b *bookRepository) DeleteBook(ctx context.Context, ID string) error {
	prefix := b.repoPrefix + ".DeleteBook"
	query := b.db.WithContext(ctx)

	if err := query.Where("id = ?", ID).Delete(&model.Book{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New(utils.ErrorDataNotFound)
		}
		return fmt.Errorf("[%s] error while execute query: %+v", prefix, err)
	}
	return nil
}
