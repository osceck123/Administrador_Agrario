package services

import (
	"github.com/osceck123/CRUD/models"
	"github.com/osceck123/CRUD/repositories"
)

// GetAllCultives obtiene todos los cultivos
func GetAllCultives() ([]models.Cultivo, error) {
	return repositories.GetAllCultives()
}

// GetCultiveByNombre obtiene el cultivo por nombre
func GetCultiveByNombre(nombre string) (models.Cultivo, error) {
	return repositories.GetCultiveByNombre(nombre)
}
