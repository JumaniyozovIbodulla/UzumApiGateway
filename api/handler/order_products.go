package handler

import (
	"net/http"
	"uzumapi/genproto/order_product_service"

	"github.com/gin-gonic/gin"
)

// CreateOrderProduct godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/orderProduct [POST]
// @Summary 	Create an orderProduct
// @Description API for creating an order products
// @Tags 		order-products
// @Accept  	json
// @Produce  	json
// @Param		order body order_product_service.CreateOrderProduct true "order-products"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) CreateOrderProduct(c *gin.Context) {
	orderProduct := order_product_service.CreateOrderProduct{}

	if err := c.ShouldBindJSON(&orderProduct); err != nil {
		handleResponse(c, h.log, "error while reading request body", http.StatusBadRequest, err)
		return
	}

	resp, err := h.grpcClient.OrderProductsService().Create(c.Request.Context(), &orderProduct)

	if err != nil {
		handleResponse(c, h.log, "error while creating order product", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.log, "Order product created successfully", http.StatusOK, resp)
}

// GetByIdOrderProduct godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/orderProduct [GET]
// @Summary 	Get an order-product
// @Description API for getting an order
// @Tags 		order-products
// @Accept  	json
// @Produce  	json
// @Param		order body order_product_service.OrderProductPrimaryKey true "order-products"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetByIdOrderProduct(c *gin.Context) {
	orderProductId := order_product_service.OrderProductPrimaryKey{}

	if err := c.ShouldBindJSON(&orderProductId); err != nil {
		handleResponse(c, h.log, "error while reading request body", http.StatusBadRequest, err)
		return
	}

	resp, err := h.grpcClient.OrderProductsService().GetById(c.Request.Context(), &orderProductId)

	if err != nil {
		handleResponse(c, h.log, "error while getting an order product", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.log, "Order product got successfully", http.StatusOK, resp)
}

// UpdateOrderProduct godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/orderProduct [PUT]
// @Summary 	Update an order product
// @Description API for update an order product
// @Tags 		order-products
// @Accept  	json
// @Produce  	json
// @Param		order body order_product_service.UpdateOrderProduct true "order-products"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) UpdateOrderProduct(c *gin.Context) {
	orderProduct := order_product_service.UpdateOrderProduct{}

	if err := c.ShouldBindJSON(&orderProduct); err != nil {
		handleResponse(c, h.log, "error while reading request body", http.StatusBadRequest, err)
		return
	}

	resp, err := h.grpcClient.OrderProductsService().Update(c.Request.Context(), &orderProduct)

	if err != nil {
		handleResponse(c, h.log, "error while updating an order product", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Order product updated successfully", http.StatusOK, resp)
}

// DeleteOrderProduct godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/orderProduct [DELETE]
// @Summary 	Delete an order product
// @Description API for delete an order product
// @Tags 		order-products
// @Accept  	json
// @Produce  	json
// @Param		order body order_product_service.OrderProductPrimaryKey true "order-products"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) DeleteOrderProduct(c *gin.Context) {
	order := order_product_service.OrderProductPrimaryKey{}

	if err := c.ShouldBindJSON(&order); err != nil {
		handleResponse(c, h.log, "error while reading request body", http.StatusBadRequest, err)
		return
	}

	resp, err := h.grpcClient.OrderProductsService().Delete(c.Request.Context(), &order)

	if err != nil {
		handleResponse(c, h.log, "error while deleting an order product", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Order product deleted successfully", http.StatusOK, resp)
}

// GetAllOrderProducts godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/orderProducts [GET]
// @Summary 	Get all orders
// @Description API for Get all order products
// @Tags 		order-products
// @Accept  	json
// @Produce  	json
// @Param		order body order_product_service.GetListOrderProductRequest true "order-products"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetAllOrderProducts(c *gin.Context) {
	orderProduct := order_product_service.GetListOrderProductRequest{}

	if err := c.ShouldBindJSON(&orderProduct); err != nil {
		handleResponse(c, h.log, "error while reading request body", http.StatusBadRequest, err)
		return
	}

	resp, err := h.grpcClient.OrderProductsService().GetAll(c.Request.Context(), &orderProduct)

	if err != nil {
		handleResponse(c, h.log, "error while getting all order products", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.log, "Order products got successfully", http.StatusOK, resp)
}