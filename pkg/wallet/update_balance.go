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

	diff := new_balance - wallet.Balance

	var transaction = &models.Transaction{}

	if diff != 0 {
		transaction.ID = uuid.New()
		transaction.Type = "credited"
		transaction.WalletID = wallet.ID
		transaction.Amount = math.Abs(diff)

		var category = &models.Category{}

		if result := h.DB.Where(&models.Category{Name: "Others"}).First(&category); result.Error != nil {
			c.AbortWithError(http.StatusInternalServerError, result.Error)
			return
		}

		transaction.CategoryID = category.ID
		transaction.Category = *category
		transaction.Recurring = false

		if diff < 0 {
			transaction.Type = "debited"
		}

		if result := h.DB.Create(&transaction); result.Error != nil {
			c.AbortWithError(http.StatusInternalServerError, result.Error)
			return
		}
	}

	wallet.Balance = new_balance

	h.DB.Save(&wallet)

	if diff != 0 {
		c.JSON(http.StatusCreated, gin.H{
			"transaction": &transaction,
			"wallet":      &wallet,
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"wallet": &wallet,
	})

}
