package country

import "time"

type Country struct {
	ID               int    `gorm:"primaryKey"`
	Name             string `gorm:"unique"`
	Code             string `gorm:"unique"`
	Continent        string
	Capital          string
	OfficialLanguage string
	Area             int `gorm:"column:area_sq_km"`
	Population       int
	Currency         string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
