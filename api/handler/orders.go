package handler

import (
	"net/http"
	"uzumapi/genproto/order_service"

	"github.com/gin-gonic/gin"
)

// CreateOrder godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/order [POST]
// @Summary 	Create an order
// @Description API for creating an order
// @Tags 		orders
// @Accept  	json
// @Produce  	json
// @Param		order body order_service.CreateOrder true "order"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) CreateOrder(c *gin.Context) {
	order := order_service.CreateOrder{}

	if err := c.ShouldBindJSON(&order); err != nil {
		handleResponse(c, h.log, "error while reading request body", http.StatusBadRequest, err)
		return
	}

	resp, err := h.grpcClient.OrderService().Create(c.Request.Context(), &order)

	if err != nil {
		handleResponse(c, h.log, "error while creating order", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.log, "Order created successfully", http.StatusOK, resp)
}

// GetByIdOrder godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/order [GET]
// @Summary 	Get an order
// @Description API for getting an order
// @Tags 		orders
// @Accept  	json
// @Produce  	json
// @Param		order body order_service.OrderPrimaryKey true "order"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetByIdOrder(c *gin.Context) {
	order := order_service.OrderPrimaryKey{}

	if err := c.ShouldBindJSON(&order); err != nil {
		handleResponse(c, h.log, "error while reading request body", http.StatusBadRequest, err)
		return
	}

	resp, err := h.grpcClient.OrderService().GetById(c.Request.Context(), &order)

	if err != nil {
		handleResponse(c, h.log, "error while getting an order", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.log, "Order got successfully", http.StatusOK, resp)
}

// UpdateOrder godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/order [PUT]
// @Summary 	Update an order
// @Description API for update an order
// @Tags 		orders
// @Accept  	json
// @Produce  	json
// @Param		order body order_service.UpdateOrder true "order"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) UpdateOrder(c *gin.Context) {
	order := order_service.UpdateOrder{}

	if err := c.ShouldBindJSON(&order); err != nil {
		handleResponse(c, h.log, "error while reading request body", http.StatusBadRequest, err)
		return
	}

	resp, err := h.grpcClient.OrderService().Update(c.Request.Context(), &order)

	if err != nil {
		handleResponse(c, h.log, "error while updating an order", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Order updated successfully", http.StatusOK, resp)
}

// DeleteOrder godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/order [DELETE]
// @Summary 	Delete an order
// @Description API for delete an order
// @Tags 		orders
// @Accept  	json
// @Produce  	json
// @Param		order body order_service.OrderPrimaryKey true "order"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) DeleteOrder(c *gin.Context) {
	order := order_service.OrderPrimaryKey{}

	if err := c.ShouldBindJSON(&order); err != nil {
		handleResponse(c, h.log, "error while reading request body", http.StatusBadRequest, err)
		return
	}

	resp, err := h.grpcClient.OrderService().Delete(c.Request.Context(), &order)

	if err != nil {
		handleResponse(c, h.log, "error while deleting an order", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Order deleted successfully", http.StatusOK, resp)
}

// GetAllOrder godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/orders [GET]
// @Summary 	Get all orders
// @Description API for Get all orders
// @Tags 		orders
// @Accept  	json
// @Produce  	json
// @Param		order body order_service.GetListOrderRequest true "order"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetAllOrder(c *gin.Context) {
	order := order_service.GetListOrderRequest{}

	if err := c.ShouldBindJSON(&order); err != nil {
		handleResponse(c, h.log, "error while reading request body", http.StatusBadRequest, err)
		return
	}

	resp, err := h.grpcClient.OrderService().GetAll(c.Request.Context(), &order)

	if err != nil {
		handleResponse(c, h.log, "error while getting all orders", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.log, "Orders got successfully", http.StatusOK, resp)
}
