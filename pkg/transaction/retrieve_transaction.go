package transaction

import (
	"net/http"
	"veil/pkg/common/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


type RetrieveTransactionQuery struct {
	ID	uuid.UUID `form:"id"`
}

func (h handler) RetrieveTransaction(c *gin.Context) {
	params := RetrieveTransactionQuery{}

	if err := c.BindQuery(&params); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var transaction =  &models.Transaction{}

	if result:= h.DB.Joins("Category").Where(&models.Transaction{ID:params.ID}).First(&transaction); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &transaction)
}
