package models

import "github.com/jinzhu/gorm"

// User representa un usuario en la base de datos
type User struct {
    gorm.Model
    FirstName   string 
    LastName    string 
    Email       string 
    PasswordHash string
}

/*
type User struct {
    gorm.Model
    FirstName   string //`json:"first_name"`
    LastName    string //`json:"last_name"`
    Email       string //`json:"email" gorm:"unique"`
    PasswordHash string //`json:"-"` // Se usa para almacenar la contraseña encriptada
}*/