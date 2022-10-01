package routes

import (
	"fmt"

	"veil/controllers/wallet.go"
	"github.com/gin-gonic/gin"
)

func Register(g *gin.RouterGroup) {
  g.GET("/lookup-wallet", RetrieveWallet)
}