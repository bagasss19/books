package feature

import (
	"net/url"
	"strconv"

	"github.com/books/service/books/model"
)

type BookPayload struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

type PropsBookUpdateByFields struct {
	Data map[string]interface{}
}

type PropsBooksGetAll struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`

	// For Pagination
	Page  int64  `json:"page"`
	Limit int64  `json:"limit"`
	Order string `json:"order"`
	Sort  string `json:"sort"`
}

type GetAllBookResponse struct {
	Books   []*model.Book
	Count   int64
	MaxPage int64
}

func (param *PropsBooksGetAll) Parse(values url.Values) error {
	param.Page = 1
	param.Limit = 5
	param.Title = ""
	param.Author = ""
	param.ISBN = ""
	param.Order = "id"
	param.Sort = "asc"

	if len(values["page"]) > 0 {
		// parse string query to int
		pg, er := strconv.ParseInt(values["page"][0], 10, 16)
		if er != nil {
			return er
		}

		// check if pg > dari query string
		if pg > int64(param.Page) {
			param.Page = int64(pg)
		}
	}

	if len(values["limit"]) > 0 {
		// parse string query to int
		li, er := strconv.ParseInt(values["limit"][0], 10, 16)
		if er != nil {
			return er
		}
		// check if limit > dari query string
		if li > int64(param.Limit) {
			param.Limit = int64(li)
		}
	}

	if len(values["order"]) > 0 && len(values["sort"]) > 0 {
		param.Order = values.Get("order")
		if values.Get("sort") == "asc" || values.Get("sort") == "desc" {
			param.Sort = values.Get("sort")
		}
	}

	if len(values.Get("title")) > 0 {
		param.Title = values.Get("title")
	}

	if len(values.Get("author")) > 0 {
		param.Author = values.Get("author")
	}

	if len(values.Get("isbn")) > 0 {
		param.ISBN = values.Get("isbn")
	}

	return nil
}
