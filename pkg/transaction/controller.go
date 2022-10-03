package transaction

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/api/transaction")
	routes.GET("/list", h.ListTransactions)
	routes.POST("/list", h.ListTransactionsWithFilters)
	routes.POST("/create", h.CreateTransaction)
	routes.PUT("/update/:id", h.UpdateTransaction)
	routes.DELETE("/delete/:id", h.DeleteTransaction)
	routes.GET("/retreive/:id", h.RetrieveTransaction)
}
