package services

import (
	"dop-backend/models"
 	"dop-backend/database"
)

// GetAllServices returns all services from the database
func GetAllServices() ([]models.Service, error) {
	var services []models.Service
	err := database.DB.Find(&services).Error
	return services, err
}


// GetServiceByID fetches a service by its ID
func GetServiceByID(id int) (models.Service, error) {
	var service models.Service
	err := database.DB.First(&service, id).Error
	return service, err
}
