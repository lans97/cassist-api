package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"slices"
)

func RequireRole(allowed_roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var userRole string
			var ok bool
			if userRole, ok = c.Get("role").(string); !ok {
				return echo.NewHTTPError(http.StatusForbidden, "Role not found in context")
			}

			if slices.Contains(allowed_roles, userRole) {
				return next(c)
			}

			return echo.NewHTTPError(http.StatusForbidden, "Insufficient permissions")
		}
	}
}
