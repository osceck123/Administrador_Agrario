package models

import "github.com/jinzhu/gorm"

type Siembra struct {
	gorm.Model

	FechaSiembra            string
	SemillaPorMetroCuadrado int
	Rendimiento             string
	Escalonamiento          string
	AbonoVerde              string
}
