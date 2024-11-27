package models

import "github.com/jinzhu/gorm"

type CosechaProductos struct {
	gorm.Model
	Rendiemiento     string
	PeriodoDeCosecha string
	FormaDeCosecha   string
	diasACosechar    string
}

type CosechaSemilla struct {
	gorm.Model
	DuracionDeSemilla string
	TIpoDeFecundacion string
	FormaDeCosecha    string
}
