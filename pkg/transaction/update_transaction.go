package transaction

import (
	"database/sql"
	"net/http"
	"veil/pkg/common/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateTransactionBody struct {
	Amount       float64      `json:"amount"`
	Type         string       `json:"type"`
	Recurring    sql.NullBool `json:"recurring"`
	RecurIntCnt  float64      `json:"recur_int_count"`
	RecurIntPer  string       `json:"recur_int_per"`
	CategoryID   uuid.UUID    `json:"category_id"`
	VoiceFile    string       `json:"voice_file"`
	VoiceMessage string       `json:"vocie_message"`
}

type UpdateTransactionParams struct {
	ID uuid.UUID `form:"id"`
}

func (h handler) UpdateTransaction(c *gin.Context) {

	updates := &UpdateTransactionBody{}

	if err := c.BindJSON(updates); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	params := UpdateTransactionParams{}

	// getting request's params
	if err := c.BindQuery(&params); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// update transaction
	var transaction = &models.Transaction{}

	if result := h.DB.Where(&models.Transaction{ID: params.ID}).First(&transaction); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	var wallet = &models.Wallet{}

	if result := h.DB.Where(&models.Wallet{ID: transaction.WalletID}).First(&wallet); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	// update wallet balance
	if updates.Type != transaction.Type {
		if transaction.Type == "debited" {
			wallet.Balance += transaction.Amount
		} else {
			wallet.Balance -= transaction.Amount
		}
	}

	if updates.Type == "debited" {
		wallet.Balance -= updates.Amount
	} else {
		wallet.Balance += updates.Amount
	}

	if updates.Recurring.Valid {
		transaction.Recurring = updates.Recurring.Bool
		if transaction.Recurring {
			transaction.RecurIntCnt = updates.RecurIntCnt
			transaction.RecurIntPer = updates.RecurIntPer
		}
	}

	if updates.VoiceMessage != "" {
		transaction.VoiceMessage = updates.VoiceMessage
		if transaction.VoiceMessage != "" {
			transaction.VoiceFile = updates.VoiceFile
		}
	}

	if (updates.CategoryID != uuid.UUID{}) {
		var category = &models.Category{}
		if result := h.DB.Where(&models.Category{ID: updates.CategoryID}).First(&category); result.Error != nil {
			c.AbortWithError(http.StatusInternalServerError, result.Error)
			return
		}
		transaction.CategoryID = updates.CategoryID
	}

	h.DB.Save(&transaction)

	// update wallet
	h.DB.Save(&wallet)

	// send updated balance and transaction struct
	c.JSON(http.StatusOK, gin.H{
		"updated_balance": wallet.Balance,
		"transaction":     &transaction,
	})

}
