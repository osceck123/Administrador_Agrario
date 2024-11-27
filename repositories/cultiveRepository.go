package repositories

import (
	utils "github.com/osceck123/CRUD/config"
	"github.com/osceck123/CRUD/models"
)

// GetAllCultives obtiene todos los cultivos
func GetAllCultives() ([]models.Cultivo, error) {
	var cultivos []models.Cultivo
	if err := utils.DB.Find(&cultivos).Error; err != nil {
		return nil, err
	}
	return cultivos, nil
}

// GetCultiveByNombre obtiene todos los usuarios
func GetCultiveByNombre(nombre string) (models.Cultivo, error) {
	var cultivo models.Cultivo
	if err := utils.DB.Where("name = ?", nombre).Find(&cultivo).Error; err != nil {
		return cultivo, err
	}
	return cultivo, nil

}

// GetCultiveByNombre obtiene todos los usuarios
func GetCultiveById(nombre string) (models.Cultivo, error) {
	var cultivo models.Cultivo
	if err := utils.DB.Where("name = ?", nombre).Find(&cultivo).Error; err != nil {
		return cultivo, err
	}
	return cultivo, nil

}
