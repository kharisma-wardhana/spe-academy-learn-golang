package auth

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kharisma-wardhana/spe-academy-learn-golang/code-with-skeleton/todo-api/config"
	"github.com/kharisma-wardhana/spe-academy-learn-golang/code-with-skeleton/todo-api/entity"
	mentity "github.com/kharisma-wardhana/spe-academy-learn-golang/code-with-skeleton/todo-api/internal/repository/mysql/entity"
)

//go:embed private_key.pem
var privateKeyBytes []byte

//go:embed public_key.pem
var publicKeyBytes []byte

type JWT struct{}

func NewJWTAuth() *JWT {
	return &JWT{}
}

type JWTAuth interface {
	GenerateToken(user *mentity.User) (string, error)
}

func (j *JWT) GenerateToken(user *mentity.User) (string, error) {

	cfg := config.NewConfig()

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return "", err
	}

	claims := &entity.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(cfg.JwtExpireDaysCount) * 24 * time.Hour)),
		},
		Email:      user.Email,
		UserID:     user.ID,
		RoleAccess: user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return fmt.Errorf("EMPTY TOKEN")
	}

	token := authHeader[7:]

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		return err
	}

	claims := &entity.Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil || !tkn.Valid {
		return err
	}

	// Set data in Local Context
	c.Locals("user_id", claims.UserID)

	return nil
}
