package models

import "github.com/jinzhu/gorm"

// User representa un usuario en la base de datos
type User struct {
	gorm.Model
	FirstName    string
	LastName     string
	Email        string
	PasswordHash string
}
