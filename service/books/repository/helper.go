package repository

import "gorm.io/gorm"

func filterBookByID(query *gorm.DB, ID string) *gorm.DB {
	return query.Where("id = ?", ID)
}

func filterBookByTitle(query *gorm.DB, title string) *gorm.DB {
	return query.Where("title = ?", title)
}

func filterBookByAuthor(query *gorm.DB, author string) *gorm.DB {
	return query.Where("author = ?", author)
}

func filterBookByISBN(query *gorm.DB, isbn string) *gorm.DB {
	return query.Where("isbn = ?", isbn)
}
