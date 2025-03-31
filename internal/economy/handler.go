package economy

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// GetEconomies handles the GET request to fetch economies
func (h *Handler) GetEconomies(c *gin.Context) {
	economies, err := h.service.GetAllEconomies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, economies)
}

