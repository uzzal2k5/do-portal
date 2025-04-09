package models

import "gorm.io/gorm"

// Service represents a DevOps service
type Service struct {
	gorm.Model
	Name   string `json:"name"`
	Status string `json:"status"`
	Logs   string `json:"logs"`
}

