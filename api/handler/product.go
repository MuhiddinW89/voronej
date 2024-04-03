package handler

import (
	"app/api/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateProduct(c *gin.Context) {

	var createProduct models.CreateProduct

	err := c.ShouldBindJSON(&createProduct)
	if err != nil {
		h.handlerResponse(c, "create product", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strorages.Product().Create(context.Background(), &createProduct)
	if err != nil {
		h.handlerResponse(c, "storage.product.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strorages.Product().GetById(context.Background(), &models.ProductPrimaryKey{ProductId: id})

	if err != nil {
		h.handlerResponse(c, "storage.product.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create product", http.StatusCreated, resp)
}
