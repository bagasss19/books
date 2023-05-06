package feature

import (
	"context"
	"errors"
	"testing"

	"github.com/books/service/books/model"
	"github.com/books/service/books/repository"
	bookMock "github.com/books/service/books/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gorm.io/gorm"
)

func Test_bookFeature_GetOneBook(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := bookMock.NewMockBookRepository(mockCtrl)

	type fields struct {
		featurePrefix string
		bookRepo      repository.BookRepository
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
		mockFunc func()
	}{
		// TODO: Add test cases.
		{
			name: "positive case",
			fields: fields{
				bookRepo: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				ID:  "uuid",
			},
			wantData: &model.Book{},
			wantErr:  false,
			mockFunc: func() {
				mockRepo.EXPECT().GetOneBook(gomock.Any(), gomock.Any()).Return(&model.Book{}, nil)
			},
		},
		{
			name: "error case - data not found",
			fields: fields{
				bookRepo: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				ID:  "uuid",
			},
			wantData: nil,
			wantErr:  true,
			mockFunc: func() {
				mockRepo.EXPECT().GetOneBook(gomock.Any(), gomock.Any()).Return(nil, gorm.ErrRecordNotFound)
			},
		},
		{
			name: "error case",
			fields: fields{
				bookRepo: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				ID:  "uuid",
			},
			wantData: nil,
			wantErr:  true,
			mockFunc: func() {
				mockRepo.EXPECT().GetOneBook(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &bookFeature{
				featurePrefix: tt.fields.featurePrefix,
				bookRepo:      tt.fields.bookRepo,
			}

			tt.mockFunc()
			gotData, err := b.GetOneBook(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("bookFeature.GetOneBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(gotData, tt.wantData, cmpopts.EquateEmpty()) {
				t.Errorf("arFeature.GetAllData() got = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func Test_bookFeature_GetAllBook(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := bookMock.NewMockBookRepository(mockCtrl)

	type fields struct {
		featurePrefix string
		bookRepo      repository.BookRepository
	}
	type args struct {
		ctx   context.Context
		props PropsBooksGetAll
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp GetAllBookResponse
		wantErr  bool
		mockFunc func()
	}{
		// TODO: Add test cases.
		{
			name: "positive case",
			fields: fields{
				bookRepo: mockRepo,
			},
			args: args{
				ctx:   context.Background(),
				props: PropsBooksGetAll{},
			},
			wantResp: GetAllBookResponse{},
			wantErr:  false,
			mockFunc: func() {
				mockRepo.EXPECT().GetAllBook(gomock.Any(), gomock.Any()).Return(repository.GetAllBookOutput{}, nil)
			},
		},
		{
			name: "negative case",
			fields: fields{
				bookRepo: mockRepo,
			},
			args: args{
				ctx:   context.Background(),
				props: PropsBooksGetAll{},
			},
			wantResp: GetAllBookResponse{},
			wantErr:  true,
			mockFunc: func() {
				mockRepo.EXPECT().GetAllBook(gomock.Any(), gomock.Any()).Return(repository.GetAllBookOutput{}, errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &bookFeature{
				featurePrefix: tt.fields.featurePrefix,
				bookRepo:      tt.fields.bookRepo,
			}
			tt.mockFunc()
			gotResp, err := b.GetAllBook(tt.args.ctx, tt.args.props)
			if (err != nil) != tt.wantErr {
				t.Errorf("bookFeature.GetAllBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(gotResp, tt.wantResp, cmpopts.EquateEmpty()) {
				t.Errorf("arFeature.GetAllData() got = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func Test_bookFeature_CreateBook(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	id := "uuid"
	defer mockCtrl.Finish()

	mockRepo := bookMock.NewMockBookRepository(mockCtrl)

	type fields struct {
		featurePrefix string
		bookRepo      repository.BookRepository
	}
	type args struct {
		ctx   context.Context
		props BookPayload
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantID   *string
		wantErr  bool
		mockFunc func()
	}{
		// TODO: Add test cases.
		{
			name: "positive case",
			fields: fields{
				bookRepo: mockRepo,
			},
			args: args{
				ctx:   context.Background(),
				props: BookPayload{},
			},
			wantID:  &id,
			wantErr: false,
			mockFunc: func() {
				mockRepo.EXPECT().CreateBook(gomock.Any(), gomock.Any()).Return(&id, nil)
			},
		},
		{
			name: "negative case",
			fields: fields{
				bookRepo: mockRepo,
			},
			args: args{
				ctx:   context.Background(),
				props: BookPayload{},
			},
			wantID:  nil,
			wantErr: true,
			mockFunc: func() {
				mockRepo.EXPECT().CreateBook(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &bookFeature{
				featurePrefix: tt.fields.featurePrefix,
				bookRepo:      tt.fields.bookRepo,
			}
			tt.mockFunc()
			gotID, err := b.CreateBook(tt.args.ctx, tt.args.props)
			if (err != nil) != tt.wantErr {
				t.Errorf("bookFeature.CreateBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotID != tt.wantID {
				t.Errorf("bookFeature.CreateBook() = %v, want %v", gotID, tt.wantID)
			}
		})
	}
}

func Test_bookFeature_UpdateBookByFields(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := bookMock.NewMockBookRepository(mockCtrl)
	type fields struct {
		featurePrefix string
		bookRepo      repository.BookRepository
	}
	type args struct {
		ctx   context.Context
		props PropsBookUpdateByFields
		ID    string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantErr  bool
		mockFunc func()
	}{
		// TODO: Add test cases.
		{
			name: "positive case",
			fields: fields{
				bookRepo: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				props: PropsBookUpdateByFields{
					Data: map[string]interface{}{
						"title":  "title",
						"isbn":   "isbn",
						"author": "author",
					},
				},
				ID: "uuid",
			},
			wantErr: false,
			mockFunc: func() {
				mockRepo.EXPECT().UpdateBook(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name: "negative case - data not found",
			fields: fields{
				bookRepo: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				props: PropsBookUpdateByFields{
					Data: map[string]interface{}{
						"title":  "title",
						"isbn":   "isbn",
						"author": "author",
					},
				},
				ID: "uuid",
			},
			wantErr: true,
			mockFunc: func() {
				mockRepo.EXPECT().UpdateBook(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(gorm.ErrRecordNotFound)
			},
		},
		{
			name: "negative case - error on update data",
			fields: fields{
				bookRepo: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				props: PropsBookUpdateByFields{
					Data: map[string]interface{}{
						"title":  "title",
						"isbn":   "isbn",
						"author": "author",
					},
				},
				ID: "uuid",
			},
			wantErr: true,
			mockFunc: func() {
				mockRepo.EXPECT().UpdateBook(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &bookFeature{
				featurePrefix: tt.fields.featurePrefix,
				bookRepo:      tt.fields.bookRepo,
			}

			tt.mockFunc()
			if err := b.UpdateBookByFields(tt.args.ctx, tt.args.props, tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("bookFeature.UpdateBookByFields() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_bookFeature_DeleteBook(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := bookMock.NewMockBookRepository(mockCtrl)

	type fields struct {
		featurePrefix string
		bookRepo      repository.BookRepository
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
		mockFunc func()
	}{
		// TODO: Add test cases.
		{
			name: "positive case",
			fields: fields{
				bookRepo: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				ID:  "uuid",
			},
			wantErr: false,
			mockFunc: func() {
				mockRepo.EXPECT().GetOneBook(gomock.Any(), gomock.Any()).Return(&model.Book{}, nil)
				mockRepo.EXPECT().DeleteBook(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name: "negative case - data not found",
			fields: fields{
				bookRepo: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				ID:  "uuid",
			},
			wantErr: true,
			mockFunc: func() {
				mockRepo.EXPECT().GetOneBook(gomock.Any(), gomock.Any()).Return(nil, gorm.ErrRecordNotFound)
			},
		},
		{
			name: "negative case - error on get data",
			fields: fields{
				bookRepo: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				ID:  "uuid",
			},
			wantErr: true,
			mockFunc: func() {
				mockRepo.EXPECT().GetOneBook(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
			},
		},
		{
			name: "error case - error on delete data",
			fields: fields{
				bookRepo: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				ID:  "uuid",
			},
			wantErr: true,
			mockFunc: func() {
				mockRepo.EXPECT().GetOneBook(gomock.Any(), gomock.Any()).Return(&model.Book{}, nil)
				mockRepo.EXPECT().DeleteBook(gomock.Any(), gomock.Any()).Return(errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &bookFeature{
				featurePrefix: tt.fields.featurePrefix,
				bookRepo:      tt.fields.bookRepo,
			}
			tt.mockFunc()
			if err := b.DeleteBook(tt.args.ctx, tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("bookFeature.DeleteBook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
