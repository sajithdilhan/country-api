package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sajithdilhan/country-api/internal/country"
	"github.com/sajithdilhan/country-api/internal/economy"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=password dbname=CountryInfo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	countryRepo := country.NewRepository(db)
	countryService := country.NewService(countryRepo)
	countryHandler := country.NewHandler(countryService)
	economyRepo := economy.NewRepository(db)
	economyService := economy.NewService(economyRepo)
	economyHandler := economy.NewHandler(economyService)

	r := gin.Default()
	r.GET("/countries", countryHandler.GetCountries)
	r.GET("/countries/id/:id", countryHandler.GetCountryById)
	r.GET("/countries/code/:code", countryHandler.GetCountryByCode)
	r.GET("/economies", economyHandler.GetEconomies)
	r.GET("/economies/id/:id", economyHandler.GetEconomyById)
	r.GET("/economies/country_id/:country_id", economyHandler.GetEconomyByCountryId)

	log.Println("Starting server on :8080...")
	r.Run(":8080")
}
