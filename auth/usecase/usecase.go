package usecase

import (
	"context"
	"crypto/sha1"
	"edittapi/application/profile"
	"fmt"
	"edittapi/models"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"edittapi/auth"
)

type AuthClaims struct {
	jwt.StandardClaims
	User *models.User `json:"user"`
}

type AuthUseCase struct {
	userRepo       auth.UserRepository
	profileUseCase profile.UseCase
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthUseCase(
	userRepo auth.UserRepository,
	profileUseCase profile.UseCase,
	hashSalt string,
	signingKey []byte,
	tokenTTLSeconds time.Duration) *AuthUseCase {
	return &AuthUseCase{
		userRepo:       userRepo,
		profileUseCase: profileUseCase,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: time.Second * tokenTTLSeconds,
	}
}

func (a *AuthUseCase) SignUp(ctx context.Context, inp auth.SignUpInput) error {
	pwd := sha1.New()
	pwd.Write([]byte(inp.Password))
	pwd.Write([]byte(a.hashSalt))

	user := &models.User{
		Email:    inp.Email,
		Username: inp.Username,
		Password: fmt.Sprintf("%x", pwd.Sum(nil)),
	}

	return a.userRepo.CreateUser(ctx, user)
}

func (a *AuthUseCase) SignIn(ctx context.Context, inp auth.SignInInput) (string, error) {
	pwd := sha1.New()
	pwd.Write([]byte(inp.Password))
	pwd.Write([]byte(a.hashSalt))
	inp.Password = fmt.Sprintf("%x", pwd.Sum(nil))

	user, err := a.userRepo.GetUser(ctx, inp.Username, inp.Password)
	if err != nil {
		return "", auth.ErrUserNotFound
	}

	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(a.expireDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(a.signingKey)
}

func (a *AuthUseCase) ParseToken(ctx context.Context, accessToken string) (*models.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return a.signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, auth.ErrInvalidAccessToken
}

func (a *AuthUseCase) AttachProfile(ctx context.Context, user *models.User) error {
	return a.profileUseCase.SetUserProfile(ctx, user)
}
