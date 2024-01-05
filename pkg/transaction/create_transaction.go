package transaction

import (
	"net/http"
	"veil/pkg/common/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateTransactionBody struct {
	WalletID     uuid.UUID `json:"wallet_id"`
	Amount       float64   `json:"amount"`
	Type         string    `json:"type"`
	Recurring    bool      `json:"recurring"`
	RecurIntCnt  float64   `json:"recur_int_count"`
	RecurIntPer  string    `json:"recur_int_per"`
	CategoryID   uuid.UUID `json:"category_id"`
	VoiceFile    string    `json:"voice_file"`
	VoiceMessage string    `json:"vocie_message"`
}

func (h handler) CreateTransaction(c *gin.Context) {
	body := &CreateTransactionBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var transaction = &models.Transaction{}

	transaction.ID = uuid.New()
	transaction.WalletID = body.WalletID
	transaction.CategoryID = body.CategoryID


	var category = &models.Category{}
	if result := h.DB.Where(&models.Category{ID:body.CategoryID}).First(&category); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	transaction.Category = *category
	transaction.Amount = body.Amount
	transaction.Type = body.Type

	transaction.Recurring = body.Recurring
	if transaction.Recurring {
		transaction.RecurIntCnt = body.RecurIntCnt
		transaction.RecurIntPer = body.RecurIntPer
	}

	transaction.VoiceMessage = body.VoiceMessage
	if transaction.VoiceMessage != "" {
		transaction.VoiceFile = body.VoiceFile
	}

	if result := h.DB.Create(&transaction); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	var wallet = &models.Wallet{}
	h.DB.Where(&models.Wallet{ID: body.WalletID}).Find(&wallet)

	if transaction.Type == "debited" {
		wallet.Balance -= transaction.Amount
	} else {
		wallet.Balance += transaction.Amount
	}

	h.DB.Save(&wallet)

	c.JSON(http.StatusCreated, gin.H{
		"transaction": &transaction,
		"wallet":      &wallet,
	})
}
