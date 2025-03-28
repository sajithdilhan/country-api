package country

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAllCountries() ([]Country, error) {
	var countries []Country
	result := r.db.Table("country_infos.countries").Find(&countries) // Adjust table name to include schema if needed
	if result.Error != nil {
		return nil, result.Error
	}
	return countries, nil
}
