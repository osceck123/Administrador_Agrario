package services

import (
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
	var user *models.User
	user.LastName = user_dto.LastName
	user.Email = user_dto.Email
	user.FirstName = user_dto.FirstName
	utils.HashPassword(user_dto.Password, 1)

	return repositories.CreateUser(user)
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
