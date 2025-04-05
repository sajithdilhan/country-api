package economy

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAllEconomies() ([]Economy, error) {
	var economies []Economy
	result := r.db.Table("country_infos.economies").Find(&economies) // Adjust table name to include schema if needed
	if result.Error != nil {
		return nil, result.Error
	}
	return economies, nil
}

func (r *Repository) GetEconomyById(id int) (*Economy, error) {
	var economy Economy
	result := r.db.Table("country_infos.economies").Where("id = ?", id).First(&economy)
	if result.Error != nil {
		return nil, result.Error
	}
	return &economy, nil
}

func (r *Repository) GetEconomyByCountryId(country_id int) (*Economy, error) {
	var economy Economy
	result := r.db.Table("country_infos.economies").Where("country_id = ?", country_id).First(&economy)
	if result.Error != nil {
		return nil, result.Error
	}
	return &economy, nil
}
