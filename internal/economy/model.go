package economy

import (
	"time"

	"github.com/sajithdilhan/country-api/internal/country"
)

type Economy struct {
	ID               uint            `gorm:"primaryKey"`
	CountryID        uint            `gorm:"not null;unique;index"`
	Country          country.Country `gorm:"foreignKey:CountryID;constraint:OnDelete:CASCADE"`
	GDPNominal       float64
	GDPPPP           float64
	InflationRate    float64
	UnemploymentRate float64
	MajorIndustries  string
	CreatedAt        time.Time
}
