package wallet

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

	routes := r.Group("/api/wallet")
	routes.POST("/retreive", h.RetrieveWallet)
	routes.POST("/create", h.CreateWallet)
	routes.PUT("/update-balance", h.UpdateBalance)
}
