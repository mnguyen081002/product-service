package route

import (
	"productservice/internal/api/middlewares"
	"productservice/internal/lib"
	"github.com/gin-gonic/gin"
)

type CmsGroupRoutes struct {
	g *gin.RouterGroup
}

func NewCmsRoutes(handler *lib.Handler, middleware *middlewares.GinMiddleware) *CmsGroupRoutes {
	g := handler.Group("/cms", middleware.Auth(true))
	return &CmsGroupRoutes{
		g: g,
	}
}
