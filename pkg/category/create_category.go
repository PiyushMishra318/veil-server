package category

import (
	"net/http"
	"veil/pkg/common/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


type CreateCategoryBody struct {
	Name	string	`json:"name"`
	Image	string	`json:"image"`
	Show	bool	`json:"show"`
}

func (h handler) CreateCategory(c *gin.Context) {
	body := &CreateCategoryBody{}

	if err:= c.BindJSON(&body); err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}

	var category = &models.Category{}

	category.ID = uuid.New()
	category.Name = body.Name
	category.Image = body.Image
	category.Show = body.Show

	if result := h.DB.Create(&category); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &category)
}