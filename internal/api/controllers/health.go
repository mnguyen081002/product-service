package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HealthController struct {
}

func NewHealthController() *HealthController {
	controller := &HealthController{}
	return controller
}

func (h *HealthController) Health(c *gin.Context) {
	Response(c, http.StatusOK, "success", map[string]interface{}{
		"status": "UP",
		"time":   time.Now().Format("2006-01-02 15:04:05"),
		"env":    viper.GetString("server.env"),
	})
}
