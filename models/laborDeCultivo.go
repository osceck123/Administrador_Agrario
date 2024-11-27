package models

import "github.com/jinzhu/gorm"

type Labor struct {
	gorm.Model

	Tarea      string
	Dificultad string
}
