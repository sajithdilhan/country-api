package economy

import (
	"net/http"
	"strconv"

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

func (h *Handler) GetEconomyById(c *gin.Context) {
	economyId := c.Param("id")
	economyIdInt, _ := strconv.Atoi(economyId)
	country, err := h.service.GetEconomyById(economyIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, country)
}

func (h *Handler) GetEconomyByCountryId(c *gin.Context) {
	countryId := c.Param("country_id")
	countryIdInt, _ := strconv.Atoi(countryId)
	country, err := h.service.GetEconomyByCountryId(countryIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, country)
}
