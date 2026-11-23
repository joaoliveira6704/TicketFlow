package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joaoliveira6704/TicketFlow/services/orders/internal/handlers"
)

func Register(r *gin.Engine, h *handlers.OrderHandler) {
	// outside /api: used by Docker/Swarm health checks
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	api := r.Group("/api")         // /api
	orders := api.Group("/orders") // /api/orders
	{
		orders.GET("", h.List)          // GET    /api/orders
		orders.GET("/:id", h.Get)       // GET    /api/orders/123
		orders.POST("", h.Create)       // POST   /api/orders
		orders.PATCH("/:id", h.Update)  // PATCH  /api/orders/123
		orders.DELETE("/:id", h.Cancel) // DELETE /api/orders/123
	}
}
