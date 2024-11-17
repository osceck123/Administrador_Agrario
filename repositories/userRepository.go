package repositories

import (
	"fmt"

	utils "github.com/osceck123/CRUD/config"
	"github.com/osceck123/CRUD/models"
)

// CreateUser crea un nuevo usuario en la base de datos
func CreateUser(user *models.User) (*models.User, error) {
	if err := utils.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByID obtiene un usuario por ID
func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := utils.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAllUsers obtiene todos los usuarios
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := utils.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// UpdateUser actualiza un usuario existente
func UpdateUser(user *models.User) (*models.User, error) {
	if err := utils.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteUser elimina un usuario por ID
func DeleteUser(id uint) error {
	if err := utils.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
func GetUserByEmail(Email string) (*models.User, error) {
	var user models.User
	fmt.Println(Email, "que tiene email")
	// Buscar por el campo Email utilizando la condición correcta
	if err := utils.DB.Where("email = ?", Email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
