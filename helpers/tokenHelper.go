package helpers

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateAllTokens creates access and refresh tokens for a user
func GenerateAllTokens(email string, firstName string, lastName string, userId string) (string, string, error) {
	accessSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	refreshSecret := os.Getenv("REFRESH_TOKEN_SECRET")
	if accessSecret == "" {
		accessSecret = "secret"
	}
	if refreshSecret == "" {
		refreshSecret = "refresh_secret"
	}

	accessClaims := jwt.MapClaims{}
	accessClaims["email"] = email
	accessClaims["first_name"] = firstName
	accessClaims["last_name"] = lastName
	accessClaims["user_id"] = userId
	accessClaims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	signedAccess, err := accessToken.SignedString([]byte(accessSecret))
	if err != nil {
		return "", "", fmt.Errorf("failed to sign access token: %w", err)
	}

	refreshClaims := jwt.MapClaims{}
	refreshClaims["user_id"] = userId
	refreshClaims["exp"] = time.Now().Add(7 * 24 * time.Hour).Unix()

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedRefresh, err := refreshToken.SignedString([]byte(refreshSecret))
	if err != nil {
		return "", "", fmt.Errorf("failed to sign refresh token: %w", err)
	}

	return signedAccess, signedRefresh, nil
}

// ValidateToken parses and validates a JWT access token string and returns the claims if valid
func ValidateToken(signedToken string) (jwt.MapClaims, error) {
	secret := os.Getenv("ACCESS_TOKEN_SECRET")
	if secret == "" {
		secret = "secret"
	}

		token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
			// Verify signing algorithm
			if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})
		if err != nil {
			return nil, err
		}
	
		if !token.Valid {
			return nil, fmt.Errorf("invalid token")
		}
	
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, fmt.Errorf("invalid token claims")
		}
	
		return claims, nil
	}