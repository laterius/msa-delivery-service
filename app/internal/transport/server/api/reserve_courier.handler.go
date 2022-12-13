package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/laterius/service_architecture_hw3/app/internal/service"
	"net/http"
)

// ReserveCourierHandler handles request to reserve courier
func ReserveCourierHandler(service service.Service) func(c *gin.Context) {
	// Request body structure
	type Body struct {
		OrderID uuid.UUID `json:"order_id"`
	}

	return func(c *gin.Context) {
		body := Body{}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
				"data":    gin.H{},
			})

			return
		}

		freeCourier, err := service.GetFreeCourier()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": err.Error(),
				"data":    gin.H{},
			})

			return
		}

		err = service.Reserve(freeCourier, body.OrderID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"data": gin.H{
					"success": false,
					"message": "could not reserve courier",
					"data": gin.H{
						"order_id": body.OrderID,
					},
				},
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"success": true,
				"message": "",
				"data": gin.H{
					"order_id": body.OrderID,
				},
			},
		})
	}
}
