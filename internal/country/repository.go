package country

import (
	"fmt"

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

func (r *Repository) GetCountryById(id int) (*Country, error) {
	var country Country
	fmt.Println("id", id)
	result := r.db.Table("country_infos.countries").Where("id = ?", id).First(&country)
	fmt.Println("country", country)
	if result.Error != nil {
		return nil, result.Error
	}
	return &country, nil
}

func (r *Repository) GetCountryByCode(code string) (*Country, error) {
	var country Country
	result := r.db.Table("country_infos.countries").Where("code = ?", code).First(&country)
	if result.Error != nil {
		return nil, result.Error
	}
	return &country, nil
}
