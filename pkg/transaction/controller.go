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
	// pagination -- untested
	routes.GET("/list", h.ListTransactions)
	// pagination and filtering -- untested
	routes.POST("/list", h.ListTransactionsWithFilters)

	// create a new transaction -- tested
	routes.POST("/create", h.CreateTransaction)

	// get transaction details -- tested
	routes.GET("/retreive/:id", h.RetrieveTransaction)

	routes.PUT("/update/:id", h.UpdateTransaction)
	routes.DELETE("/delete/:id", h.DeleteTransaction)
}
