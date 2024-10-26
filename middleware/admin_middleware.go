package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"github.com/golang-jwt/jwt/v4"
)

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		user := c.Get("user")
		if user == nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Unauthorized",
			})
		}

		claims := user.(*jwt.Token).Claims.(jwt.MapClaims)

		isAdmin, ok := claims["isAdmin"].(bool)
		if !ok || !isAdmin {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": "You must be an admin",
			})
		}

		return next(c)
	}
}
