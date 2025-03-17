package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lans97/cassist-api/internal/database"
	"github.com/lans97/cassist-api/internal/firebase"
	"github.com/lans97/cassist-api/internal/models"
)

func FirebaseAuth(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        authHeader := c.Request().Header.Get("Authorization")

        if authHeader == "" {
            return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
        }

        tokenString := authHeader[len("Bearer "):]

        token, err := firebase.AuthClient.VerifyIDToken(c.Request().Context(), tokenString)
        if err != nil {
            return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
        }

        var user models.User
        res := database.DB.Model(&models.User{}).Where("uuid = ?", token.UID).Preload("Role").First(&user)
        if res.Error != nil {
            userRecord, err := firebase.AuthClient.GetUser(c.Request().Context(), token.UID)
            if err != nil {
                return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retreive user from details")
            }

            ev := true

            user = models.User{
                UUID: userRecord.UID,
                Email: userRecord.Email,
                DisplayName: userRecord.DisplayName,
                EmailVerified: &ev,
                RoleID: 2,
            }
            res := database.DB.Create(&user)
            if res.Error != nil {
                return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create new internal user")
            }
        }

        c.Set("google_uid", user.UUID)
        c.Set("user_id", user.ID)
        c.Set("role", user.Role.Name)

        return next(c)
    }
}
