package models

import "time"

// ContactForm defines the structure for a contact form submission
type ContactForm struct {
    Name    string    `json:"name"`
    Email   string    `json:"email"`
    Mobile  string    `json:"mobile"`
    Message string    `json:"message"`
    CreatedAt time.Time `json:"created_at"`
}
