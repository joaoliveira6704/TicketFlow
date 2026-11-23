package models

import "time"

type OrderStatus string

const (
	OrderPending   OrderStatus = "pending"
	OrderConfirmed OrderStatus = "confirmed"
	OrderCancelled OrderStatus = "cancelled"
	OrderExpired   OrderStatus = "expired"
)

type Order struct {
	ID        string      `db:"id"         json:"id"`
	UserID    string      `db:"user_id"    json:"userId"`
	EventID   string      `db:"event_id"   json:"eventId"`
	Status    OrderStatus `db:"status"     json:"status"`
	Total     int64       `db:"total"      json:"total"` // cents, never float for money
	CreatedAt time.Time   `db:"created_at" json:"createdAt"`
	ExpiresAt time.Time   `db:"expires_at" json:"expiresAt"`
	Items     []OrderItem `db:"-"          json:"items"`
}

type OrderItem struct {
	ID           string `db:"id"             json:"id"`
	OrderID      string `db:"order_id"       json:"-"`
	TicketTypeID string `db:"ticket_type_id" json:"ticketTypeId"`
	Quantity     int    `db:"quantity"       json:"quantity"`
	UnitPrice    int64  `db:"unit_price"     json:"unitPrice"`
}
