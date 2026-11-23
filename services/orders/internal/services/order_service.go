package services

import (
	"context"

	"github.com/joaoliveira6704/TicketFlow/services/orders/internal/models"
)

type OrderService struct {
	model *models.OrderModel
}

func NewOrderService(m *models.OrderModel) *OrderService {
	return &OrderService{model: m}
}

func (s *OrderService) List(ctx context.Context) ([]models.Order, error) {
	return s.model.FindAll(ctx)
}

func (s *OrderService) GetByID(ctx context.Context, id string) (*models.Order, error) {
	return s.model.FindByID(ctx, id)
}
