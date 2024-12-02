package system

import "github.com/gin-gonic/gin"

type SystemHandler struct{}

// Health HealthCheck godoc
// @Summary 死活監視用
// @Tags health-check
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Router /v1/health [get]
func (h *SystemHandler) Health(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"Status": "ok",
	})
}

func NewSystemHandler() *SystemHandler {
	return &SystemHandler{}
}
