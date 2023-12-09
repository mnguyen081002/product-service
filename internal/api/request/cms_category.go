package request

import uuid "github.com/satori/go.uuid"

type CreateCategoryRequest struct {
	Name     string     `json:"name" binding:"required"`
	ParentID *uuid.UUID `json:"parent_id"`
}

type UpdateCategoryRequest struct {
	Name     string     `json:"name" binding:"required"`
	ParentID *uuid.UUID `json:"parent_id"`
}

type ListCategoryRequest struct {
	PageOptions
	ParentID *uuid.UUID `form:"parent_id" json:"parent_id"`
	IsParent bool       `form:"is_parent" json:"is_parent"`
}
