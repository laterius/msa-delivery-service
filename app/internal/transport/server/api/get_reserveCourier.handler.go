package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	s "github.com/laterius/service_architecture_hw3/app/internal/service"
	"net/http"
)

// GetReserveCourierHandler handles request to get reserve courier by orderID
func GetReserveCourierHandler(service s.Service) func(c *gin.Context) {
	return func(c *gin.Context) {

		id := c.Param("orderId")

		orderId, err := uuid.Parse(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "order id can't parse",
				"data":    gin.H{},
			})
			return
		}

		reservation, err := service.Get(orderId)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "reserve courier not found",
				"data":    gin.H{},
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Reserve courier found",
			"data": gin.H{
				"order_id": reservation.OrderId,
			},
		})
	}
}
