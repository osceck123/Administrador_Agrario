package models

import "github.com/jinzhu/gorm"

type Huerta struct {
	gorm.Model

	Ubicacion              string
	DistanciaEntreCultivos string
	DistanciaEntreLineas   string
	AsociarCon             string
	RotarCon               string
	Espacio                string
	Sombra                 bool
	cultivoEnRecipiente    bool
}
