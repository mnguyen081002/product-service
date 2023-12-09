package main

import (
	"context"
	"fmt"
	"os"
	"productservice/internal/api/request"
	"productservice/internal/models"
	"reflect"
	"testing"

	uuid "github.com/satori/go.uuid"
)

func Test_cmsCategoryService_CreateCategory(t *testing.T) {
	parentCat, err := cmsCategoryService.CreateCategory(ctx, request.CreateCategoryRequest{
		Name: "Parent Category",
	})
	if err != nil {
		t.Errorf("cmsCategoryService.CreateCategory() error = %v", err)
		return
	}

	invalidUUID := uuid.NewV4()

	type args struct {
		ctx context.Context
		req request.CreateCategoryRequest
	}
	tests := []struct {
		name    string
		args    args
		wantCat *models.Category
		wantErr bool
	}{
		{
			name: "Valid category",
			args: args{
				ctx: ctx,
				req: request.CreateCategoryRequest{
					Name: "Test Category",
				},
			},
			wantCat: &models.Category{
				Name:     "Test Category",
				ParentID: nil,  // No parent ID
				NoSub:    true, // Default value
			},
			wantErr: false,
		},
		{
			name: "Missing name",
			args: args{
				ctx: ctx,
				req: request.CreateCategoryRequest{},
			},
			wantCat: nil,
			wantErr: true,
		},
		{
			name: "Empty name",
			args: args{
				ctx: ctx,
				req: request.CreateCategoryRequest{
					Name: "",
				},
			},
			wantCat: nil,
			wantErr: true,
		},
		{
			name: "Valid category with parent ID",
			args: args{
				ctx: ctx,
				req: request.CreateCategoryRequest{
					Name:     "Subcategory",
					ParentID: &parentCat.ID,
				},
			},
			wantCat: &models.Category{
				Name:     "Subcategory",
				ParentID: &parentCat.ID,
				NoSub:    true,
			},
			wantErr: false,
		},
		{
			name: "Invalid parent ID",
			args: args{
				ctx: ctx,
				req: request.CreateCategoryRequest{
					Name:     "Subcategory",
					ParentID: &invalidUUID,
				},
			},
			wantCat: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotCat, err := cmsCategoryService.CreateCategory(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("cmsCategoryService.CreateCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantCat != nil {
				if !reflect.DeepEqual(gotCat.Name, tt.wantCat.Name) {
					t.Errorf("cmsCategoryService.CreateCategory() = %v, want %v", gotCat.Name, tt.wantCat.Name)
				}
				if !reflect.DeepEqual(gotCat.ParentID, tt.wantCat.ParentID) {
					t.Errorf("cmsCategoryService.CreateCategory() = %v, want %v", gotCat.ParentID, tt.wantCat.ParentID)
				}
				if !reflect.DeepEqual(gotCat.NoSub, tt.wantCat.NoSub) {
					t.Errorf("cmsCategoryService.CreateCategory() = %v, want %v", gotCat.NoSub, tt.wantCat.NoSub)
				}
			}
		})
	}
}

func Test_cmsCategoryService_UpdateCategory(t *testing.T) {
	invalidUUID := uuid.NewV4()

	type args struct {
		ctx context.Context
		req request.UpdateCategoryRequest
	}
	tests := []struct {
		name          string
		args          args
		wantCat       *models.Category
		wantParentCat *models.Category
		wantErr       bool
		withParent    bool
	}{
		{
			name: "Valid category",
			args: args{
				ctx: ctx,
				req: request.UpdateCategoryRequest{
					Name: "Test Category",
				},
			},
			wantCat: &models.Category{
				Name:  "Test Category",
				NoSub: true,
			},
			wantErr: false,
		},
		{
			name: "Missing name",
			args: args{
				ctx: ctx,
				req: request.UpdateCategoryRequest{},
			},
			wantCat: nil,
			wantErr: true,
		},
		{
			name: "Empty name",
			args: args{
				ctx: ctx,
				req: request.UpdateCategoryRequest{
					Name: "",
				},
			},
			wantCat: nil,
			wantErr: true,
		},
		{
			name: "Valid category with parent ID",
			args: args{
				ctx: ctx,
				req: request.UpdateCategoryRequest{
					Name: "Subcategory", // can them parent id
				},
			},
			wantCat: &models.Category{
				Name:  "Subcategory", // can them parent id
				NoSub: true,
			},
			wantParentCat: &models.Category{
				NoSub: false,
			},
			withParent: true,
			wantErr:    false,
		},
		{
			name: "Invalid parent ID",
			args: args{
				ctx: ctx,
				req: request.UpdateCategoryRequest{
					Name:     "Subcategory",
					ParentID: &invalidUUID,
				},
			},
			wantCat: nil,
			wantErr: true,
		},
		{
			name: "Valid category with no parent ID",
			args: args{
				ctx: ctx,
				req: request.UpdateCategoryRequest{
					Name: "Subcategory",
				},
			},
		},
		{
			name: "Valid update category has parent has one sub category to no parent ID",
			args: args{
				ctx: ctx,
				req: request.UpdateCategoryRequest{
					Name: "Subcategory",
				},
			},
			wantParentCat: &models.Category{
				NoSub:    true,
				ParentID: nil,
			},
			wantCat: &models.Category{
				Name:     "Subcategory",
				NoSub:    true,
				ParentID: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parentCat := &models.Category{}
			reqCreate := request.CreateCategoryRequest{
				Name: "Subcategory",
			}
			var err error
			if tt.wantParentCat != nil {
				parentCat, err = cmsCategoryService.CreateCategory(ctx, request.CreateCategoryRequest{
					Name: "Parent Category",
				})
				reqCreate.ParentID = &parentCat.ID

				if err != nil {
					t.Errorf("cmsCategoryService.CreateCategory() error = %v", err)
					return
				}
			}

			if tt.withParent {
				tt.args.req.ParentID = &parentCat.ID
				tt.wantCat.ParentID = &parentCat.ID
			}

			cat, err := cmsCategoryService.CreateCategory(ctx, reqCreate)
			if err != nil {
				t.Errorf("cmsCategoryService.CreateCategory() error = %v", err)
				return
			}

			gotCat, err := cmsCategoryService.UpdateCategory(tt.args.ctx, tt.args.req, cat.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("cmsCategoryService.UpdateCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantCat != nil {
				if !reflect.DeepEqual(gotCat.Name, tt.wantCat.Name) {
					t.Errorf("cmsCategoryService.UpdateCategory() = %v, want %v", gotCat.Name, tt.wantCat.Name)
				}
				if !reflect.DeepEqual(gotCat.ParentID, tt.wantCat.ParentID) {
					t.Errorf("cmsCategoryService.UpdateCategory() = %v, want %v", gotCat.ParentID, tt.wantCat.ParentID)
				}
				if !reflect.DeepEqual(gotCat.NoSub, tt.wantCat.NoSub) {
					t.Errorf("cmsCategoryService.UpdateCategory() = %v, want %v", gotCat.NoSub, tt.wantCat.NoSub)
				}
			}

			if tt.wantParentCat != nil {
				parentCat, err := ufw.CategoryRepository.FindByID(&db, ctx, parentCat.ID)
				if err != nil {
					t.Errorf("cmsCategoryService.UpdateCategory() error = %v", err)
					return
				}
				if !reflect.DeepEqual(parentCat.NoSub, tt.wantParentCat.NoSub) {
					t.Errorf("cmsCategoryService.UpdateCategory() = %v, want %v", parentCat.NoSub, tt.wantParentCat.NoSub)
				}
			}
		})
	}
}

func TestMain(m *testing.M) {
	SetUp()
	fmt.Println("Start test")
	os.Exit(m.Run())
}
