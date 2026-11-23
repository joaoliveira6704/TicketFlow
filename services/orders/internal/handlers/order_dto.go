package handlers

type CreateOrderRequest struct {
	EventID string            `json:"eventId" binding:"required,uuid"`
	Items   []CreateOrderItem `json:"items"   binding:"required,min=1,dive"`
}

type CreateOrderItem struct {
	TicketTypeID string `json:"ticketTypeId" binding:"required,uuid"`
	Quantity     int    `json:"quantity"     binding:"required,min=1,max=10"`
}
