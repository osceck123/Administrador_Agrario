package utils

import (
	"fmt"
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
	defer DB.Close()

	// Migración de modelos
	DB.AutoMigrate(
		&models.User{},
		&models.CosechaProductos{},
		&models.CosechaSemilla{},
		&models.Cultivo{},
		&models.Huerta{},
		&models.Labor{},
		&models.Patio{},
		&models.Siembra{},
	)

	// Insertar datos de prueba
	// Cosecha de Producto
	cosechaProducto := models.CosechaProductos{
		Rendiemiento:     "2 a 3 kg por metro cuadrado",
		PeriodoDeCosecha: "todo los meses",
		FormaDeCosecha:   "Hoja por hoja",
	}
	DB.Create(&cosechaProducto)

	// Cosecha de Semilla
	cosechaSemilla := models.CosechaSemilla{
		DuracionDeSemilla: "4",
		TIpoDeFecundacion: "Cruzada",
		FormaDeCosecha:    "Corte de tallos florales que se van secando",
	}
	DB.Create(&cosechaSemilla)

	// Labor
	laborDeCultivo := models.Labor{
		Tarea:      "Riego y deshierbe",
		Dificultad: "Fácil",
	}
	DB.Create(&laborDeCultivo)

	// Huerta
	huertaCultivo := models.Huerta{
		Ubicacion:              "Canteros",
		DistanciaEntreCultivos: "12 a 20 centimetros",
		DistanciaEntreLineas:   "50 a 70 centimetros",
		AsociarCon:             "Maíz/Lechuga/Escarola",
		RotarCon:               "Legumbres",
		Espacio:                "poco",
		Sombra:                 false,
	}
	DB.Create(&huertaCultivo)

	// Siembra
	siembraCultivo := models.Siembra{
		Rendimiento:             "100 semillas = 50 plantines",
		FechaSiembra:            "todos los meses, menos enero y julio",
		SemillaPorMetroCuadrado: 12,
		Escalonamiento:          "1 siembra mensual",
		AbonoVerde:              "no",
	}
	DB.Create(&siembraCultivo)

	fmt.Println("Datos de prueba insertados correctamente")
}
