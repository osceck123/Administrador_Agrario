package main

import (
	"github.com/gin-gonic/gin"
	utils "github.com/osceck123/CRUD/config"
	"github.com/osceck123/CRUD/controllers"
	"github.com/osceck123/CRUD/middleware"
)

func main() {
	// Inicializar la base de datos
	utils.InitDB()

	// Inicializar el router de Gin
	router := gin.Default()

	// Definir rutas
	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.GetAllUsers)
	router.GET("/users/:id", controllers.GetUserByID)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)

	// Rutas de autenticación
	router.POST("/login", controllers.Login)
	router.GET("/logout", controllers.Logout)

	// Rutas protegidas por el JWT middleware
	protected := router.Group("/")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.GET("/profile", controllers.GetUserByID) // Ejemplo de ruta protegida
	}

	// Iniciar el servidor
	router.Run(":8080")
}
