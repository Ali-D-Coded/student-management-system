package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtHandler struct {
	secretAccess  string
	secretRefresh string
}

func NewJwtHandler() *JwtHandler {
	secretAccess := os.Getenv("JWT_AT_SECRET")
	if secretAccess == "" {
		secretAccess = "khfshdfgdshfgdsfhsdgfh"
	}

	secretRefresh := os.Getenv("JWT_RT_SECRET")
	if secretRefresh == "" {
		secretRefresh = "rlgdkflfgjkrldgjkjdfgj"
	}

	return &JwtHandler{
		secretAccess:  secretAccess,
		secretRefresh: secretRefresh,
	}
}

type Token struct {
	Access        string `json:"access"`
	Refresh       string `json:"refresh"`
	AccessExpires int64  `json:"accessExpires"`
}

func (h *JwtHandler) GenerateToken(payload map[string]interface{}) (Token, error) {
    fmt.Printf("Secret: %s\n", h.secretAccess)

	// Calculate expiration time for the access token in seconds
    accessExpires := time.Now().Add(time.Minute * 30).Unix() 

    // Create the access token
    accessToken := jwt.New(jwt.SigningMethodHS256)
    claims := accessToken.Claims.(jwt.MapClaims)
    for key, value := range payload {
        claims[key] = value
    }
    claims["exp"] = accessExpires // 1 minute expiration in seconds

    access, err := accessToken.SignedString([]byte(h.secretAccess))
    if err != nil {
        return Token{}, err
    }

    // Create the refresh token
    refreshToken := jwt.New(jwt.SigningMethodHS256)
    refreshClaims := refreshToken.Claims.(jwt.MapClaims)
    for key, value := range payload {
        refreshClaims[key] = value
    }
    refreshClaims["exp"] = time.Now().Add(time.Hour * 72).Unix() // 3 days expiration in seconds

    refresh, err := refreshToken.SignedString([]byte(h.secretRefresh))
    if err != nil {
        return Token{}, err
    }

    

    // Return the Token struct with appropriate values
    return Token{
        Access:        access,
        Refresh:       refresh,
        AccessExpires: accessExpires,
    }, nil
}



func (h *JwtHandler) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(h.secretAccess), nil
	})

	fmt.Println( err)

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}