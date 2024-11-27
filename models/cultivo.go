package models

import (
	"github.com/jinzhu/gorm"
)

type Cultivo struct {
	gorm.Model
	Nombre             string
	Familia            string
	Comportamiento     string
	Caractresticas     string
	SemillasPorGramo   string
	CosechaProductosID *uint
	CosechaSemillaID   *uint
	HuertasID          *uint
	LaboresID          *uint
	SiembrasID         *uint

	CosechaProductos []CosechaProductos `gorm:"foreignKey:CosechaProductosID"` // Relación uno a muchos
	CosechaSemilla   []CosechaSemilla   `gorm:"foreignKey:CosechaSemillaID"`
	Huertas          []Huerta           `gorm:"foreignKey:HuertasID"`
	Labores          []Labor            `gorm:"foreignKey:LaboresID"`
	Siembras         []Siembra          `gorm:"foreignKey:SiembrasID;"`
}
