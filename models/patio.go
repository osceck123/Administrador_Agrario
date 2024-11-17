package models

import "github.com/jinzhu/gorm"

type Patio struct {
	gorm.Model
	Imagen   string // Ruta a la imagen subida
	Anchura  float64
	Altura   float64
	Sectores []Sector
}

type Sector struct {
	gorm.Model
	PatioID     uint
	Cultivo     string
	Coordenadas []float64 // Array de coordenadas (x, y)
}
