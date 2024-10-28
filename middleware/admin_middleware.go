package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("In admin middleware")

		user := c.Get("user")
		if user == nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Unauthorized",
			})
		}

		// Directly assert user to jwt.MapClaims
		claims, ok := user.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Invalid token claims",
			})
		}

		isAdmin, ok := claims["isAdmin"].(bool)
		if !ok || !isAdmin {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": "You must be an admin",
			})
		}

		return next(c)
	}
}
