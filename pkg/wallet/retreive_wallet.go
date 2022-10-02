package wallet

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "veil/pkg/common/models"
)


type RetrieveWalletBody struct {
	Phone	string	`json:"phone"`
}

func (h handler) RetrieveWallet(c *gin.Context) {
    body := RetrieveWalletBody{}

    // getting request's body
    if err := c.BindJSON(&body); err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }

	phone := body.Phone


	var wallet = &models.Wallet{}

	if result := h.DB.Where(&models.Wallet{Phone: phone}).First(&wallet); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

    c.JSON(http.StatusCreated, &wallet)
}