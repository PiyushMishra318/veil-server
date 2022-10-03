package wallet

import (
	"net/http"
	"veil/pkg/common/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateWalletBody struct {
	Phone       string  `json:"phone"`
	OwnerName   string  `json:"owner_name"`
	Expenditure float64 `json:"expenditure"`
	Savings     float64 `json:"savings"`
	Balance     float64 `json:"balance"`
}

func (h handler) CreateWallet(c *gin.Context) {
	body := CreateWalletBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// create wallet
	var wallet = &models.Wallet{}
	wallet.Balance = 0

	wallet.Expenditure = body.Expenditure
	wallet.Savings = body.Savings
	wallet.OwnerName = body.OwnerName
	wallet.Phone = body.Phone
	wallet.ID = uuid.New()

	if result := h.DB.Create(&wallet); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &wallet)
}
