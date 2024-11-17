package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Cultivo struct {
	gorm.Model
	Nombre       string
	Tipo         string
	Sector       int
	FechaSiembra time.Time
	Estado       string
}
