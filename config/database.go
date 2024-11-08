package utils

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/osceck123/CRUD/models"
)

var DB *gorm.DB

// InitDB inicializa la conexión con la base de datos
func InitDB() {
	var err error
	// Conexión con PostgreSQL
	DB, err = gorm.Open("postgres", "host=localhost port=5433 user=admin dbname=campo password=usuario_admin sslmode=disable")
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}
	// Migración de modelos
	DB.AutoMigrate(&models.User{})
}
