package country

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler struct
type Handler struct {
	service *Service
}

// NewHandler initializes a new Handler
func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

// GetCountries handles the GET request to fetch all countries
func (h *Handler) GetCountries(c *gin.Context) {
	countries, err := h.service.GetAllCountries()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, countries)
}
