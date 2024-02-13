package handler

import (
	"net/http"
	"strconv"

	"github.com/MrSaveliy/store/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createCategory(c *gin.Context) {
	var input models.Category
	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Category.CreateCategory(input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getCategories(c *gin.Context) {

}

func (h *Handler) getCategoryById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "invalid id param")
		return
	}

	category, err := h.services.Category.GetById(id)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"name":        category.Name,
		"descriptiom": category.Description,
	})
}

func (h *Handler) updateCategory(c *gin.Context) {

}

func (h *Handler) deleteCategory(c *gin.Context) {

}
