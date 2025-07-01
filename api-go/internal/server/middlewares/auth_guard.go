package middlewares

import (
	"log"
	"net/http"

	"github.com/Gabriel-Schiestl/api-go/internal/infra/database"
	"github.com/Gabriel-Schiestl/api-go/internal/infra/database/connection"
	"github.com/Gabriel-Schiestl/api-go/internal/infra/mappers"
	"github.com/Gabriel-Schiestl/api-go/internal/infra/ports"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	service := ports.NewJWTService()

	return func(c *gin.Context) {
		if c.FullPath() == "/auth/login" || c.FullPath() == "/auth/logout" || (c.FullPath() == "/users/" && c.Request.Method == "POST") {
			c.Next()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not provided"})
			c.Abort()
			return
		}

		// Extract token from "Bearer <token>" format
		var authToken string
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			authToken = authHeader[7:]
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		claims, err := service.ExtractClaims(authToken)
		if err != nil {
			log.Printf("Error extracting claims: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		log.Printf("JWT Claims: %v", claims)
		userID := claims["sub"].(string)
		log.Printf("Extracted user ID: %s", userID)

		user, err := database.NewUserRepository(connection.Db, mappers.UserMapper{}).FindById(userID)
		if err != nil {
			log.Printf("Error finding user by ID %s: %v", userID, err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		log.Printf("Found user: %s (ID: %s)", user.GetName(), user.GetID())
		c.Set("userID", user.GetID())
		c.Next()
	}
}