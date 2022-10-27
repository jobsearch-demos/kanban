package auth

import (
	"encoding/base64"
	errors "jobsearch-demos/kanban/pkg/error"
	"time"

	"github.com/golang-jwt/jwt"
)

// Create jwt token from map
func CreateJWTToken(data map[string]interface{}, jwtSecret string, ttl int) (string, error) {
	// ttl is time to live in seconds
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Second * time.Duration(ttl)).Unix(),
	}
	for k, v := range data {
		claims[k] = v
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Encode token with secret
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func DecodeJWTToken(token string, jwtSecret string) (map[string]interface{}, error) {
	// Parse token with secret
	parsedToken, err := jwt.Parse(string(token), func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if !parsedToken.Valid {
		return nil, errors.CustomError{Code: 401, Message: "Invalid token"}
	}
	return parsedToken.Claims.(jwt.MapClaims), nil
}

func CreateEncodedJWTToken(data map[string]interface{}, jwtSecret string) (string, error) {
	token, err := CreateJWTToken(data, jwtSecret, 3600)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString([]byte(token)), nil
}

func DecodeEncodedJWTToken(token string, jwtSecret string) (map[string]interface{}, error) {
	// Decode base64 token
	decodedToken, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}
	return DecodeJWTToken(string(decodedToken), jwtSecret)
}
