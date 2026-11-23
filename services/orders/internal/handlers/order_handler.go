package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/joaoliveira6704/TicketFlow/services/orders/internal/services"
)

type OrderHandler struct {
	service *services.OrderService
}

func NewOrderHandler(s *services.OrderService) *OrderHandler {
	return &OrderHandler{service: s}
}

func (h *OrderHandler) List(c *gin.Context) {
	orders, err := h.service.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL", "message": err.Error()}})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) Get(c *gin.Context) {
	order, err := h.service.GetByID(c.Request.Context(), c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": "NOT_FOUND", "message": "Order not found"}})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) Create(c *gin.Context) {
	var req CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "VALIDATION", "message": err.Error()}})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "not implemented yet"})
}

func (h *OrderHandler) Update(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented yet"})
}

func (h *OrderHandler) Cancel(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
