package middleware

import (
    "net/http"
    "strings"
    "github.com/osceck123/CRUD/utils"
    "github.com/gin-gonic/gin"
)

// JWTAuthMiddleware es un middleware para verificar la autenticación JWT
func JWTAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Obtener el token de los headers Authorization
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
            c.Abort()
            return
        }

        // El formato debe ser "Bearer <token>"
        parts := strings.SplitN(authHeader, " ", 2)
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
            c.Abort()
            return
        }

        // Validar el token
        token, err := utils.ValidateJWT(parts[1])
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
            c.Abort()
            return
        }

        // Almacenar el token validado en el contexto para usarlo en el controlador
        c.Set("token", token)
        c.Next()
    }
}
