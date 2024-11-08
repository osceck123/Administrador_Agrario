package controllers

import (
	"net/http"
	"time"

	"github.com/osceck123/CRUD/dto"
	"github.com/osceck123/CRUD/services"
	"github.com/osceck123/CRUD/utils"

	"github.com/gin-gonic/gin"
)

// Login maneja la autenticación de usuario
func Login(c *gin.Context) {
	var loginDTO dto.LoginDTO
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Verificar las credenciales
	user, err := services.GetUserByEmail(loginDTO.Email)
	if err != nil || !utils.CheckPasswordHash(loginDTO.Password, user.PasswordHash) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generar JWT
	token, err := utils.GenerateJWT(user.Email, 24*time.Hour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating JWT"})
		return
	}

	// Crear el UserDTO para enviar los datos del usuario al frontend
	userDTO := dto.UserDTO{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	// Responder con el token y el DTO del usuario
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  userDTO,
	})
}

// Logout maneja el cierre de sesión (básicamente solo elimina el token en el cliente)
func Logout(c *gin.Context) {
	// No hay una lógica del lado del servidor para el logout en JWT
	// El cliente simplemente debe eliminar el token
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
