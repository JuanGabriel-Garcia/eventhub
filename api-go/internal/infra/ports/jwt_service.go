package ports

import (
	"fmt"
	"os"
	"time"

	"github.com/Gabriel-Schiestl/api-go/internal/domain/services"
	"github.com/golang-jwt/jwt/v5"
)

type jwtService struct {
	secretKey []byte
}

func NewJWTService() services.IJWTService {
	return &jwtService{
		secretKey: []byte(os.Getenv("JWT_SECRET_KEY")),
	}
}

func (s *jwtService) GenerateToken(userID string) (*string, error) {
	claims := jwt.MapClaims{
        "sub":  userID, 
        "iat":  time.Now().Unix(),
        "exp":  time.Now().Add(time.Hour * 24).Unix(),
    }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    tokenString, err := token.SignedString(s.secretKey)
    if err != nil {
        fmt.Println("Erro ao criar o token:", err)
        return nil, fmt.Errorf("error creating token: %w", err)
    }

	return &tokenString, nil
}

func (s *jwtService) ExtractClaims(token string) (map[string]interface{}, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

