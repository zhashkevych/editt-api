package auth

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type AuthClaims struct {
	jwt.StandardClaims
	username string
	password string
}

type Authorizer struct {
	username       string
	passwordHash   string
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthorizer(username, pHash, hashSalt string, signingKey []byte, expireDuration time.Duration) *Authorizer {
	return &Authorizer{
		username:       username,
		passwordHash:   pHash,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: expireDuration * time.Second,
	}
}

func (a *Authorizer) GenerateToken(username, password string) (string, error) {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	if a.username == username && a.passwordHash == password {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &AuthClaims{
			username: username,
			password: password,
		})

		return token.SignedString(a.signingKey)
	}

	return "", ErrInvalidCredentials
}

func (a *Authorizer) ParseToken(accessToken string) error {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return a.signingKey, nil
	})

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return nil
	}

	return ErrInvalidAccessToken
}

// Authorizer HTTP middleware
func (a *Authorizer) Middleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if headerParts[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	err := a.ParseToken(headerParts[1])
	if err != nil {
		status := http.StatusBadRequest
		if err == ErrInvalidAccessToken {
			status = http.StatusUnauthorized
		}

		c.AbortWithStatus(status)
		return
	}
}
