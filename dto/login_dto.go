// dto/login_dto.go
package dto

// LoginDTO es el objeto que usará el frontend para enviar las credenciales de usuario
type LoginDTO struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}
