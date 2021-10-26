package middleware

import (
	"AltaStore/config"

	"github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(config.GetConfig().JwtSecretKey),
	})
}

func ExtractToken(ctx echo.Context) (string, error) {
	user := ctx.Get("user").(*jwt.Token)
	if user.Valid {
		claim := user.Claims.(jwt.MapClaims)
		userId := claim["userId"].(string)
		return userId, nil
	}
	return "", nil
}
