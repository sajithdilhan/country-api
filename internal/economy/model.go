package economy

import (
	"time"
)

type Economy struct {
	ID               uint `gorm:"primaryKey"`
	CountryID        uint `gorm:"not null;unique;index"`
	GDPNominal       float64
	GDPPPP           float64 `gorm:"column:gdp_ppp"`
	InflationRate    float64
	UnemploymentRate float64
	MajorIndustries  string
	CreatedAt        time.Time
}
