package services

type IJWTService interface {
	GenerateToken(userID string) (*string, error)
	ExtractClaims(token string) (map[string]interface{}, error)
}