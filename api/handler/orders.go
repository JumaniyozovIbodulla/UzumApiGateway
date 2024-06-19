package handler

import (
	"net/http"
	"uzumapi/genproto/order_service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
// @Router 		/api/v1/order/{id} [GET]
// @Summary 	Get an order
// @Description API for getting an order
// @Tags 		orders
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetByIdOrder(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating orderId", http.StatusBadRequest, err.Error())
		return
	}

	order := order_service.OrderPrimaryKey{
		Id: id,
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
// @Router 		/api/v1/order/{id} [DELETE]
// @Summary 	Delete an order
// @Description API for delete an order
// @Tags 		orders
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating orderId", http.StatusBadRequest, err.Error())
		return
	}

	order := order_service.OrderPrimaryKey{
		Id: id,
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
// @Param   	search query string false "search"
// @Param    	page query int false "page"
// @Param    	limit query int false "limit"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetAllOrder(c *gin.Context) {
	search := c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponse(c, h.log, "error while parsing page", http.StatusBadRequest, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)

	if err != nil {
		handleResponse(c, h.log, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}

	order := order_service.GetListOrderRequest{
		Offset: (page -1) * limit,
		Limit: limit,
		SearchByCustomerName: search,
	}

	resp, err := h.grpcClient.OrderService().GetAll(c.Request.Context(), &order)

	if err != nil {
		handleResponse(c, h.log, "error while getting all orders", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.log, "Orders got successfully", http.StatusOK, resp)
}
