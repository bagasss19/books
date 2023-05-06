package repository

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/books/service/books/model"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// func Test_bookRepository_GetAllBook(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	assert.NoError(t, err)

// 	dialector := postgres.New(postgres.Config{
// 		DSN:                  "sqlmock_db_0",
// 		DriverName:           "postgres",
// 		Conn:                 db,
// 		PreferSimpleProtocol: true,
// 	})
// 	gormDB, err := gorm.Open(dialector, &gorm.Config{})
// 	assert.NoError(t, err)

// 	query := `SELECT * FROM "books" ORDER BY $1 $2`
// 	queryCount := `SELECT count(*) FROM "books"`

// 	type fields struct {
// 		repoPrefix string
// 		db         *gorm.DB
// 	}
// 	type args struct {
// 		ctx   context.Context
// 		props PropsBooksGetAll
// 	}
// 	tests := []struct {
// 		name     string
// 		fields   fields
// 		args     args
// 		wantResp GetAllBookOutput
// 		wantErr  bool
// 		mockFunc func(args)
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "positive case",
// 			fields: fields{
// 				db: gormDB,
// 			},
// 			args: args{
// 				ctx: context.Background(),
// 				props: PropsBooksGetAll{
// 					Order: "id",
// 					Sort:  "asc",
// 				},
// 			},
// 			wantErr: false,
// 			wantResp: GetAllBookOutput{
// 				Data: []*model.Book{
// 					{
// 						ID:    "3ed329b1-4e32-4d4a-99dd-7b1d41b6b354",
// 						Title: "title",
// 					},
// 				},
// 			},
// 			mockFunc: func(args args) {
// 				mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(args.props.Order, args.props.Sort).WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).AddRow("3ed329b1-4e32-4d4a-99dd-7b1d41b6b354", "title"))
// 				mock.ExpectQuery(regexp.QuoteMeta(queryCount)).WithArgs().WillReturnRows(sqlmock.NewRows([]string{"max"}).AddRow(5))
// 				mock.ExpectCommit()
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			b := &bookRepository{
// 				repoPrefix: tt.fields.repoPrefix,
// 				db:         tt.fields.db,
// 			}

// 			tt.mockFunc(tt.args)
// 			gotResp, err := b.GetAllBook(tt.args.ctx, tt.args.props)
// 			if err != nil {
// 				t.Errorf("Failed to meet expectations, got error: %v", err)
// 			}

// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("bookRepository.GetAllBook() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}

// 			if !cmp.Equal(gotResp, tt.wantResp, cmpopts.EquateEmpty()) {
// 				t.Errorf("arFeature.GetAllData() got = %v, want %v", gotResp, tt.wantResp)
// 			}
// 		})
// 	}
// }

func Test_bookRepository_GetOneBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)

	query := `SELECT * FROM "books" WHERE id = $1`
	type fields struct {
		repoPrefix string
		db         *gorm.DB
	}
	type args struct {
		ctx context.Context
		ID  string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantData *model.Book
		wantErr  bool
		mockFunc func(args)
	}{
		// TODO: Add test cases.
		{
			name: "positive case",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx: context.Background(),
				ID:  "uuid",
			},
			wantErr: false,
			wantData: &model.Book{
				ID:    "uuid",
				Title: "title",
			},
			mockFunc: func(args args) {
				mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(args.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).AddRow("uuid", "title"))
			},
		},
		{
			name: "error case - error not found",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx: context.Background(),
				ID:  "uuid",
			},
			wantErr:  true,
			wantData: nil,
			mockFunc: func(args args) {
				mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(args.ID).WillReturnError(gorm.ErrRecordNotFound)
			},
		},
		{
			name: "error case - error while exec query",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx: context.Background(),
				ID:  "uuid",
			},
			wantErr:  true,
			wantData: nil,
			mockFunc: func(args args) {
				mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(args.ID).WillReturnError(errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &bookRepository{
				repoPrefix: tt.fields.repoPrefix,
				db:         tt.fields.db,
			}

			tt.mockFunc(tt.args)
			gotData, err := b.GetOneBook(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("bookRepository.GetOneBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(gotData, tt.wantData, cmpopts.EquateEmpty()) {
				t.Errorf("arFeature.GetAllData() got = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func Test_bookRepository_CreateBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)

	query := `INSERT INTO "books" ("id","title","author","isbn","created_time","updated_time") VALUES ($1,$2,$3,$4,$5,$6)`
	ID := "uuid"
	now := time.Now()
	type fields struct {
		repoPrefix string
		db         *gorm.DB
	}
	type args struct {
		ctx   context.Context
		props model.Book
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantID   *string
		wantErr  bool
		mockFunc func(args)
	}{
		// TODO: Add test cases.
		{
			name: "positive case",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx: context.Background(),
				props: model.Book{
					ID:          "uuid",
					Title:       "title",
					Author:      "author",
					ISBN:        "isbn",
					CreatedTime: now,
					UpdatedTime: now,
				},
			},
			wantErr: false,
			wantID:  &ID,
			mockFunc: func(args args) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(
					args.props.ID,
					args.props.Title,
					args.props.Author,
					args.props.ISBN,
					args.props.CreatedTime,
					args.props.UpdatedTime).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "error case",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx: context.Background(),
				props: model.Book{
					ID:          "uuid",
					Title:       "title",
					Author:      "author",
					ISBN:        "isbn",
					CreatedTime: now,
					UpdatedTime: now,
				},
			},
			wantErr: true,
			wantID:  nil,
			mockFunc: func(args args) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(
					args.props.ID,
					args.props.Title,
					args.props.Author,
					args.props.ISBN,
					args.props.CreatedTime,
					args.props.UpdatedTime).
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &bookRepository{
				repoPrefix: tt.fields.repoPrefix,
				db:         tt.fields.db,
			}

			tt.mockFunc(tt.args)
			gotID, err := b.CreateBook(tt.args.ctx, tt.args.props)
			if (err != nil) != tt.wantErr {
				t.Errorf("bookRepository.CreateBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.EqualValues(t, gotID, tt.wantID)
		})
	}
}

func Test_bookRepository_DeleteBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)

	query := `DELETE FROM "books" WHERE id = $1`
	ID := "uuid"

	type fields struct {
		repoPrefix string
		db         *gorm.DB
	}
	type args struct {
		ctx context.Context
		ID  string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantErr  bool
		mockFunc func(args)
	}{
		// TODO: Add test cases.
		{
			name: "positive case",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx: context.Background(),
				ID:  ID,
			},
			wantErr: false,
			mockFunc: func(args args) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(args.ID).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "error case - data not found",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx: context.Background(),
				ID:  ID,
			},
			wantErr: true,
			mockFunc: func(args args) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(args.ID).
					WillReturnError(gorm.ErrRecordNotFound)
				mock.ExpectCommit()
			},
		},
		{
			name: "error case - error while exec query",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx: context.Background(),
				ID:  ID,
			},
			wantErr: true,
			mockFunc: func(args args) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(args.ID).
					WillReturnError(errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &bookRepository{
				repoPrefix: tt.fields.repoPrefix,
				db:         tt.fields.db,
			}

			tt.mockFunc(tt.args)
			if err := b.DeleteBook(tt.args.ctx, tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("bookRepository.DeleteBook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_bookRepository_UpdateBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)

	query := `UPDATE "books" SET "title"=$1 WHERE id = $2`
	ID := "uuid"

	type fields struct {
		repoPrefix string
		db         *gorm.DB
	}
	type args struct {
		ctx    context.Context
		props  model.Book
		fields []string
		ID     string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantErr  bool
		mockFunc func(args)
	}{
		// TODO: Add test cases.
		{
			name: "positive case",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx: context.Background(),
				ID:  ID,
				props: model.Book{
					Title: "title",
				},
			},
			wantErr: false,
			mockFunc: func(args args) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(args.props.Title, args.ID).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "error case - data not found",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx: context.Background(),
				ID:  ID,
				props: model.Book{
					Title: "title",
				},
			},
			wantErr: true,
			mockFunc: func(args args) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(args.props.Title, args.ID).
					WillReturnError(gorm.ErrRecordNotFound)
				mock.ExpectCommit()
			},
		},
		{
			name: "error case - error on exec query",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx: context.Background(),
				ID:  ID,
				props: model.Book{
					Title: "title",
				},
			},
			wantErr: true,
			mockFunc: func(args args) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(args.props.Title, args.ID).
					WillReturnError(errors.New("error"))
				mock.ExpectCommit()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &bookRepository{
				repoPrefix: tt.fields.repoPrefix,
				db:         tt.fields.db,
			}

			tt.mockFunc(tt.args)
			if err := b.UpdateBook(tt.args.ctx, tt.args.props, tt.args.fields, tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("bookRepository.UpdateBook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
