package transaction

import (
	"net/http"
	"strings"
	"veil/pkg/common/models"
	"veil/pkg/common/types"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ListTransactionsParams struct {
	Limit int `form:"limit"`
}

func (h handler) ListTransactions(c *gin.Context) {

	params := ListTransactionsParams{}

	// getting request's params
	if err := c.BindQuery(&params); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	limit := params.Limit

	var transactions []models.Transaction

	if result := h.DB.Limit(limit).Find(&transactions); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &transactions)
}

func (h handler) ListTransactionsWithFilters(c *gin.Context) {

	body := types.Pagination{}

	// getting request's params
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	limit := body.Limit
	offset := body.Offset
	sort := body.Sort
	attributes := strings.Split(body.Select, ",")
	filter := body.Filter

	var transactions []models.Transaction

	queryDB := h.DB.Where(filter).Session(&gorm.Session{})

	if offset > 0 {
		queryDB.Offset(offset)
	}

	if limit > 0 {
		queryDB.Limit(limit)
	}

	if sort != "" {
		queryDB.Order(sort)
	}

	if len(attributes) > 0 {
		queryDB.Select(attributes)
	}

	if result := queryDB.Find(&transactions); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusFound, &transactions)
}
