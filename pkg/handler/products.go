package handler

import (
	"net/http"

	"github.com/MrSaveliy/store/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createProduct(c *gin.Context) {
	var input models.Product
	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Product.CreateProduct(input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getProductById(c *gin.Context) {

}

func (h *Handler) updateProduct(c *gin.Context) {

}

func (h *Handler) deleteProduct(c *gin.Context) {

}
