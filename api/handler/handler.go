package handler

import (
	"app/config"
	"app/storage"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	cfg       *config.Config
	strorages storage.StorageI
}

type Response struct {
	Status      int
	Description string
	Data        interface{}
}

func NewHandler(cfg *config.Config, storages storage.StorageI) *Handler {
	return &Handler{
		cfg:       cfg,
		strorages: storages,
	}
}

func (h *Handler) handlerResponse(c *gin.Context, path string, code int, message interface{}) {
	response := Response{
		Status:      code,
		Description: path,
		Data:        message,
	}

	c.JSON(code, response)
}
