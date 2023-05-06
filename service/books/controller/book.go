package controller

import (
	"encoding/json"
	"net/http"

	"github.com/books/service/books/feature"
	"github.com/books/utils"
	"github.com/go-chi/chi/v5"
)

// Get Book
// @Summary      Get Book List
// @Description  show list of Book
// @Tags         Book
// @Param        title   query      string  false  "Books Title"
// @Param        author   query      string  false  "Books Author"
// @Param        isbn   query      string  false  "Books ISBN"
// @Param        page   query      string  false  "Page. Default is 1"
// @Param        limit   query      string  false  "Limit. Default is 5"
// @Param        order   query      string  false  "Order data by field"
// @Param        sort   query      string  false  "Sorting. ASC or DESC"
// @Router       /books [get]
func (b bookController) GetAllBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req feature.PropsBooksGetAll

	err := req.Parse(r.URL.Query())
	if err != nil {
		b.SendBadRequest(w, err.Error())
		return
	}
	resp, err := b.bookFeature.GetAllBook(ctx, req)
	if err != nil {
		b.RespondWithJSON(w, 500, err.Error(), nil)
		return
	}

	b.RespondWithJSON(w, 200, utils.OK, resp)
}

// Get One Book
// @Summary      Get one Book
// @Description  show Book by ID
// @Tags         Book
// @Param        id   path      string  true  "book id"
// @Router       /books/{id} [get]
func (b bookController) GetOneBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	resp, err := b.bookFeature.GetOneBook(ctx, id)
	if err != nil {
		if err.Error() == utils.ErrorDataNotFound {
			b.RespondWithJSON(w, 404, utils.ErrorDataNotFound, nil)
			return
		}
		b.RespondWithJSON(w, 500, err.Error(), nil)
		return
	}

	b.RespondWithJSON(w, 200, utils.OK, resp)
}

// Create Book
// @Summary      Create Book
// @Description  Create Book
// @Tags         Book
// @Param        payload    body   feature.BookPayload  true  "body payload"
// @Router       /books [post]
func (b bookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var request feature.BookPayload

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := b.bookFeature.CreateBook(ctx, request)
	if err != nil {
		b.RespondWithJSON(w, 500, err.Error(), nil)
		return
	}

	b.RespondWithJSON(w, 200, "Create Book Success!", id)
}

// Delete Book
// @Summary      Delete Book
// @Description  Delete Book by ID
// @Tags         Book
// @Param        id   path      string  true  "Book ID"
// @Router       /books/{id} [delete]
func (b bookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	err := b.bookFeature.DeleteBook(ctx, id)
	if err != nil {
		if err.Error() == utils.ErrorDataNotFound {
			b.RespondWithJSON(w, 404, utils.ErrorDataNotFound, nil)
			return
		}
		b.RespondWithJSON(w, 500, err.Error(), nil)
		return
	}

	b.RespondWithJSON(w, 200, "Delete Book Success!!", nil)
}

type BookUpdateRequestJson struct {
	Data map[string]interface{} `json:"data"`
}

// Update Book
// @Summary      Update Book
// @Description  Update Book with dynamic fields. editable field : isbn, author, title
// @Tags         Book
// @Param        id   path    string  true  "Book ID"
// @Param        payload    body   BookUpdateRequestJson  true  "enter desired field that want to update"
// @Router       /books/{id} [put]
func (b bookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var (
		request BookUpdateRequestJson
		id      = chi.URLParam(r, "id")
	)

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedBook := feature.PropsBookUpdateByFields{
		Data: request.Data,
	}

	err = b.bookFeature.UpdateBookByFields(ctx, updatedBook, id)
	if err != nil {
		if err.Error() == utils.ErrorDataNotFound {
			b.RespondWithJSON(w, 404, utils.ErrorDataNotFound, nil)
			return
		}
		b.RespondWithJSON(w, 500, err.Error(), nil)
		return
	}

	b.RespondWithJSON(w, 200, "Update Book Success!!", updatedBook)
}
