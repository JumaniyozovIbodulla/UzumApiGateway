package handler

import (
	"net/http"
	"uzumapi/genproto/order_product_service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
// @Router 		/api/v1/orderProduct/{id} [GET]
// @Summary 	Get an order-product
// @Description API for getting an order
// @Tags 		order-products
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetByIdOrderProduct(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating order productId", http.StatusBadRequest, err.Error())
		return
	}

	orderProductId := order_product_service.OrderProductPrimaryKey{
		Id: id,
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
// @Router 		/api/v1/orderProduct/{id} [DELETE]
// @Summary 	Delete an order product
// @Description API for delete an order product
// @Tags 		order-products
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) DeleteOrderProduct(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating order productId", http.StatusBadRequest, err.Error())
		return
	}

	order := order_product_service.OrderProductPrimaryKey{
		Id: id,
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
// @Param    	page query int false "page"
// @Param    	limit query int false "limit"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetAllOrderProducts(c *gin.Context) {
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

	orderProduct := order_product_service.GetListOrderProductRequest{
		Offset: (page - 1) * limit,
		Limit:  limit,
	}

	resp, err := h.grpcClient.OrderProductsService().GetAll(c.Request.Context(), &orderProduct)

	if err != nil {
		handleResponse(c, h.log, "error while getting all order products", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.log, "Order products got successfully", http.StatusOK, resp)
}
