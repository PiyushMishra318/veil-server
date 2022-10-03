package wallet

import (
	"math"
	"net/http"
	"veil/pkg/common/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateBalanceBody struct {
	ID         uuid.UUID `json:"id"`
	NewBalance float64   `json:"new_balance"`
}

func (h handler) UpdateBalance(c *gin.Context) {
	body := UpdateBalanceBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// create wallet
	var wallet = &models.Wallet{}

	new_balance := body.NewBalance

	if result := h.DB.Find(&models.Wallet{ID: body.ID}).First(&wallet); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	diff := wallet.Balance - new_balance

	var transaction = &models.Transaction{}

	transaction.Type = "credited"
	transaction.WalletID = wallet.ID
	transaction.Amount = math.Abs(diff)
	transaction.CategoryID = uuid.New()
	transaction.Recurring = false

	if diff < 0 {
		transaction.Type = "debited"
	}

	if result := h.DB.Create(&transaction); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"transaction": &transaction,
		"wallet":      &wallet,
	})
}
