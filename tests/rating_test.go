package main

import (
	"context"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	config2 "productservice/config"
	"productservice/internal/api/request"
	"productservice/internal/api_errors"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/models"
	"productservice/internal/repository"
	"productservice/internal/services"
	"reflect"
	"testing"
)

func Test_ratingService_CreateRating(t *testing.T) {
	type fields struct {
		db            infrastructure.Database
		dbTransaction infrastructure.DatabaseTransaction
		ratingService domain.RatingService
		ufw           *repository.UnitOfWork
		config        *config2.Config
		logger        *zap.Logger
	}
	type args struct {
		ctx context.Context
		req request.CreateRating
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantR   *models.Rating
		wantErr error
	}{
		{
			name: "Valid rating",
			fields: fields{
				db:            db,
				dbTransaction: dbTransaction,
				ufw:           ufw,
				config:        config,
				logger:        logger,
			},
			args: args{
				ctx: ctx,
				req: request.CreateRating{
					Rate:    5,
					Comment: "Test comment",
					Images:  []string{"https://test.com/image1.jpg", "https://test.com/image2.jpg"},
				},
			},
		},
		{
			name: "Missing product ID",
			fields: fields{
				db:            db,
				dbTransaction: dbTransaction,
				ufw:           ufw,
				config:        config,
				logger:        logger,
			},
			args: args{
				ctx: ctx,
				req: request.CreateRating{
					Rate:    5,
					Comment: "Test comment",
					Images:  []string{"https://test.com/image1.jpg", "https://test.com/image2.jpg"},
				},
			},
			wantErr: errors.New(api_errors.ErrInvalidProductID),
		},
		{
			name: "Invalid product ID",
			fields: fields{
				db:            db,
				dbTransaction: dbTransaction,
				ufw:           ufw,
				config:        config,
				logger:        logger,
			},
			args: args{
				ctx: ctx,
				req: request.CreateRating{
					ProductID: "invalid-uuid",
					Rate:      5,
					Comment:   "Test comment",
					Images:    []string{"https://test.com/image1.jpg", "https://test.com/image2.jpg"},
				},
			},
			wantErr: errors.New(api_errors.ErrInvalidProductID),
		},
	}
	for _, tt := range tests {
		cat, err := cmsCategoryService.CreateCategory(ctx, request.CreateCategoryRequest{
			Name: "Test Category",
		})
		if err != nil {
			t.Errorf("cmsCategoryService.CreateCategory() error = %v", err)
			return
		}

		product, err := cmsProductService.CreateProduct(ctx, request.CreateProductRequest{
			Name:        "Test Product",
			Description: "Test Description",
			Price:       100000,
			Quantity:    0,
			Images:      nil,
			CategoryID:  cat.ID.String(),
		})
		if err != nil {
			t.Errorf("cmsProductService.CreateProduct() error = %v", err)
			return
		}

		if tt.wantErr == nil {
			raterId := uuid.FromStringOrNil(userId)
			tt.wantR = &models.Rating{
				ProductID: &product.ID,
				RaterID:   &raterId,
				Rating:    tt.args.req.Rate,
				Comment:   tt.args.req.Comment,
				Images:    tt.args.req.Images,
			}

			tt.args.req.ProductID = product.ID.String()
		}

		t.Run(tt.name, func(t *testing.T) {
			s := service.NewRatingService(tt.fields.db, tt.fields.dbTransaction, tt.fields.ufw, tt.fields.config, tt.fields.logger)
			gotR, err := s.CreateRating(tt.args.ctx, tt.args.req)
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("CreateRating() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if (err != nil) && (err.Error() != tt.wantErr.Error()) {
				t.Errorf("CreateRating() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantR != nil {
				if !reflect.DeepEqual(gotR.ProductID, tt.wantR.ProductID) {
					t.Errorf("CreateRating() = %v, want %v", gotR.ProductID, tt.wantR.ProductID)
				}
				if !reflect.DeepEqual(gotR.RaterID, tt.wantR.RaterID) {
					t.Errorf("CreateRating() = %v, want %v", gotR.RaterID, tt.wantR.RaterID)
				}
				if !reflect.DeepEqual(gotR.Rating, tt.wantR.Rating) {
					t.Errorf("CreateRating() = %v, want %v", gotR.Rating, tt.wantR.Rating)
				}
				if !reflect.DeepEqual(gotR.Comment, tt.wantR.Comment) {
					t.Errorf("CreateRating() = %v, want %v", gotR.Comment, tt.wantR.Comment)
				}
				if !reflect.DeepEqual(gotR.Images, tt.wantR.Images) {
					t.Errorf("CreateRating() = %v, want %v", gotR.Images, tt.wantR.Images)
				}

				product, err := ufw.ProductRepository.GetById(&db, ctx, product.ID.String())
				if err != nil {
					t.Errorf("CreateRating() error = %v", err)
					return
				}

				if product.MedRating != float64(tt.wantR.Rating) {
					t.Errorf("CreateRating() = %v, want %v", product.MedRating, tt.wantR.Rating)
				}

				if product.RatingCount != 1 {
					t.Errorf("CreateRating() = %v, want %v", product.RatingCount, 1)
				}
			}
		})
	}
}
