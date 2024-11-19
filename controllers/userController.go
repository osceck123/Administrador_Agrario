package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/osceck123/CRUD/dto"
	"github.com/osceck123/CRUD/models"
	"github.com/osceck123/CRUD/services"
)

// CreateUser crea un nuevo usuario
func CreateUser(c *gin.Context) {
	var user dto.CreateUserDTO
	fmt.Print(&user, `usuario`)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := services.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdUser)
}

// GetUserByID obtiene un usuario por ID
func GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := services.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetAllUsers obtiene todos los usuarios
func GetAllUsers(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener usuarios"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// UpdateUser actualiza un usuario
func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = uint(id)
	updatedUser, err := services.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser elimina un usuario por ID
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := services.DeleteUser(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el usuario"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado"})
}

// GetUserProfile obtiene la información del perfil del usuario
func GetUserProfile(c *gin.Context) {
	// Obtenemos el email del token JWT
	token, _ := c.Get("token")
	claims := token.(*jwt.Token).Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	user, err := services.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Mapear el usuario al DTO
	userDTO := dto.UserDTO{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	// Responder con los datos del usuario
	c.JSON(http.StatusOK, gin.H{"user": userDTO})
}
