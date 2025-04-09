// models/portfolio.go
package models

import "time"

// Portfolio struct to represent portfolio content
type Portfolio struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	SkillSet    string    `json:"skill_set"`
	Experience  string    `json:"experience"`
	CreatedAt   time.Time `json:"created_at"`
}
