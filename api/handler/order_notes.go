package handler

import (
	"net/http"
	"uzumapi/genproto/order_notes"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateOrderNotes godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/orderNote [POST]
// @Summary 	Create an order notes
// @Description API for creating an order notes
// @Tags 		order-notes
// @Accept  	json
// @Produce  	json
// @Param		order body order_notes.CreateOrderNotes true "order-notes"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) CreateOrderNotes(c *gin.Context) {
	orderNotes := order_notes.CreateOrderNotes{}

	if err := c.ShouldBindJSON(&orderNotes); err != nil {
		handleResponse(c, h.log, "error while reading request body", http.StatusBadRequest, err)
		return
	}

	resp, err := h.grpcClient.OrderProductNotesService().Create(c.Request.Context(), &orderNotes)

	if err != nil {
		handleResponse(c, h.log, "error while creating order notes", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.log, "Order notes created successfully", http.StatusOK, resp)
}

// GetByIdOrderNotes godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/orderNote/{id} [GET]
// @Summary 	Get an order notes
// @Description API for getting an order notes
// @Tags 		order-notes
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetByIdOrderNotes(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating order notes id", http.StatusBadRequest, err.Error())
		return
	}

	orderNotes := order_notes.OrderNotesPrimaryKey{
		Id: id,
	}

	resp, err := h.grpcClient.OrderProductNotesService().GetById(c.Request.Context(), &orderNotes)

	if err != nil {
		handleResponse(c, h.log, "error while getting an order notes", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.log, "Order notes got successfully", http.StatusOK, resp)
}

// UpdateOrderNotes godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/orderNote [PUT]
// @Summary 	Update an order notes
// @Description API for update an order notes
// @Tags 		order-notes
// @Accept  	json
// @Produce  	json
// @Param		order body order_notes.UpdateOrderNotes true "order-notes"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) UpdateOrderNotes(c *gin.Context) {
	orderNotes := order_notes.UpdateOrderNotes{}

	if err := c.ShouldBindJSON(&orderNotes); err != nil {
		handleResponse(c, h.log, "error while reading request body", http.StatusBadRequest, err)
		return
	}

	resp, err := h.grpcClient.OrderProductNotesService().Update(c.Request.Context(), &orderNotes)

	if err != nil {
		handleResponse(c, h.log, "error while updating an order notes", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Order notes updated successfully", http.StatusOK, resp)
}

// DeleteOrderNotes godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/orderNote/{id} [DELETE]
// @Summary 	Delete an order notes
// @Description API for delete an order notes
// @Tags 		order-notes
// @Accept  	json
// @Produce  	json
// @Param		id path string true "id"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) DeleteOrderNotes(c *gin.Context) {
	id := c.Param("id")

	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.log, "error while validating order notes id", http.StatusBadRequest, err.Error())
		return
	}

	orderNotesPk := order_notes.OrderNotesPrimaryKey{
		Id: id,
	}

	resp, err := h.grpcClient.OrderProductNotesService().Delete(c.Request.Context(), &orderNotesPk)

	if err != nil {
		handleResponse(c, h.log, "error while deleting an order notes", http.StatusBadRequest, err.Error())
		return
	}
	handleResponse(c, h.log, "Order notes deleted successfully", http.StatusOK, resp)
}

// GetAllOrderNotes godoc
// @Security ApiKeyAuth
// @Router 		/api/v1/orderNotes [GET]
// @Summary 	Get all order notes
// @Description API for Get all order notes
// @Tags 		order-notes
// @Accept  	json
// @Produce  	json
// @Param    	page query int false "page"
// @Param    	limit query int false "limit"
// @Success		200  {object} models.Response
// @Failure		400  {object} models.Response
// @Failure 	404  {object} models.Response
// @Failure 	500  {object} models.Response
func (h *handler) GetAllOrderNotes(c *gin.Context) {
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

	orderNotes := order_notes.GetListOrderNotesRequest{
		Offset: (page - 1) * limit,
		Limit:  limit,
	}

	resp, err := h.grpcClient.OrderProductNotesService().GetAll(c.Request.Context(), &orderNotes)

	if err != nil {
		handleResponse(c, h.log, "error while getting all order notes", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.log, "Order notes got successfully", http.StatusOK, resp)
}
