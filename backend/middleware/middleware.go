package middleware

import (
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/palSagnik/daily-expenses-application/config"
	"github.com/palSagnik/daily-expenses-application/models"
)



func VerifyToken() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(config.SESSION_SECRET),
			JWTAlg: jwtware.HS256,
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "failure", "message": "invalid or expired JWT"})
		},
	})
}

func GenerateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"userid": user.UserID,
		"email": user.Email,
		"expiry": time.Now().Add(time.Hour * time.Duration(config.SESSION_EXPIRY)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.SESSION_SECRET))
	if err != nil {
		return "", err
	}
	
	return t, nil
}