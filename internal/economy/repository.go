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
