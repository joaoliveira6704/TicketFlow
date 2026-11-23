package models

import (
	"context"
	"errors"
)

var ErrNotFound = errors.New("not found")

type OrderModel struct{}

func NewOrderModel() *OrderModel { return &OrderModel{} }

func (m *OrderModel) FindAll(ctx context.Context) ([]Order, error) {
	return []Order{}, nil
}

func (m *OrderModel) FindByID(ctx context.Context, id string) (*Order, error) {
	return nil, ErrNotFound
}
