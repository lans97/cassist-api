package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lans97/cassist-api/internal/models"
)

func CustomErrorHandler(err error, c echo.Context) {
    var status_code int
    var message string

    if httpErr, ok := err.(*echo.HTTPError); ok {
        status_code = httpErr.Code
        message = httpErr.Message.(string)
    } else {
        status_code = http.StatusInternalServerError
        message = "Internal Server Error"
    }

    c.JSON(status_code, models.APIResponse {
        StatusCode: status_code,
        Message: message,
    })
}
