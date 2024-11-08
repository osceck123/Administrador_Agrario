package utils

import (
    "errors"
    "fmt"
    "golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt/v5"
    "time"
)

// ErrWeakPassword indicates that the provided password doesn't meet security requirements.
var ErrWeakPassword = errors.New("password must be at least 8 characters long and contain a mix of uppercase, lowercase, numbers, and symbols")

// ErrMalformedJWT signifies that the provided token is invalid or has an incorrect format.
var ErrInvalidToken = errors.New("invalid JWT token")

// JwtSecretKey should be replaced with a strong, randomly generated key for production use (see https://godoc.org/golang.org/x/crypto/rand)
var JwtSecretKey = []byte("replace_with_strong_key") // **IMPORTANT: Replace with a secure key**

// HashPassword encrypts a password using bcrypt with a configurable cost for increased security.
func HashPassword(password string, cost int) (string, error) {
    if len(password) < 8 {
        return "", ErrWeakPassword
    }

    // Adjust cost based on processing power and security needs (higher cost is slower but more secure)
    if cost < bcrypt.MinCost || cost > bcrypt.MaxCost {
        return "", fmt.Errorf("invalid cost: %d (minimum: %d, maximum: %d)", cost, bcrypt.MinCost, bcrypt.MaxCost)
    }

    bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
    return string(bytes), err
}

// CheckPasswordHash compares a password with a hashed password.
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

// GenerateJWT generates a JWT token with a configurable expiration time.
func GenerateJWT(email string, expiration time.Duration) (string, error) {
    claims := jwt.MapClaims{
        "email": email,
        "exp":   time.Now().Add(expiration).Unix(), // Adjust expiration as needed
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(JwtSecretKey)
}

// ValidateJWT validates a JWT token and returns its claims.
func ValidateJWT(tokenString string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Verificamos el método de firma del token (Spanish for "Verify token's signing method")
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, ErrInvalidToken
        }
        return JwtSecretKey, nil
    })
}