package utils

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"productservice/internal/api_errors"

	uuid "github.com/satori/go.uuid"
)

func GetUserUUIDFromContext(ctx context.Context) (uuid.UUID, error) {
	sid := ctx.Value("x-user-id").(string)

	u, err := uuid.FromString(sid)
	if err != nil {
		return uuid.Nil, errors.New(api_errors.ErrInvalidUserID)
	}

	return u, nil
}

func GetStoreIDFromContext(ctx context.Context) string {
	return ctx.Value("x-store-id").(string)
}

func GetUUIDFromParam(c *gin.Context, param string) (uuid.UUID, error) {
	sid := c.Param(param)

	u, err := uuid.FromString(sid)
	if err != nil {
		return uuid.Nil, err
	}
	return u, nil
}

func GetPageCount(total int64, limit int64) int {
	if total == 0 {
		return 0
	}

	if total%limit != 0 {
		return int(total/limit + 1)
	}

	return int(total / limit)
}
