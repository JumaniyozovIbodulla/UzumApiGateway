package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /v1/ondemand-order [post]
// @Summary Create On Demand Order
// @Description API for creating on demand order
// @Tags order
// @Accept  json
// @Produce  json
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handler) GetListCategory(c *gin.Context) {

	c.JSON(http.StatusOK, "")
}
