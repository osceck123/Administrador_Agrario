package services

import (
	"fmt"

	"github.com/osceck123/CRUD/dto"
	"github.com/osceck123/CRUD/models"
	"github.com/osceck123/CRUD/repositories"
	"github.com/osceck123/CRUD/utils"
)

func ConvertToUserDTO(user models.User) dto.UserDTO {
	return dto.UserDTO{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}

// CreateUser crea un usuario
func CreateUser(user_dto *dto.CreateUserDTO) (*models.User, error) {
	// Inicializar la estructura de usuario
	user := &models.User{
		LastName:  user_dto.LastName,
		Email:     user_dto.Email,
		FirstName: user_dto.FirstName,
	}

	// Hashear la contraseña
	pass, err := utils.HashPassword(user_dto.Password, 4)
	if err != nil {
		return nil, fmt.Errorf("error al hashear la contraseña: %w", err)
	}

	// Asignar el hash de la contraseña
	user.PasswordHash = pass

	// Crear el usuario en el repositorio
	createdUser, err := repositories.CreateUser(user)
	if err != nil {
		return nil, fmt.Errorf("error al crear el usuario: %w", err)
	}

	return createdUser, nil
}

// GetUserByID obtiene un usuario por ID
func GetUserByID(id uint) (*models.User, error) {
	return repositories.GetUserByID(id)
}

// GetAllUsers obtiene todos los usuarios
func GetAllUsers() ([]models.User, error) {
	return repositories.GetAllUsers()
}

// UpdateUser actualiza un usuario
func UpdateUser(user *models.User) (*models.User, error) {
	return repositories.UpdateUser(user)
}

// DeleteUser elimina un usuario
func DeleteUser(id uint) error {
	return repositories.DeleteUser(id)
}

// funcion de encontar por mail
func GetUserByEmail(email string) (*models.User, error) {
	return repositories.GetUserByEmail(email)
}
