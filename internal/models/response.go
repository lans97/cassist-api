package models

import "github.com/labstack/echo/v4"

type APIResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
}

func JSONResponse(c echo.Context, status int, message string, data ...any) error {
    return c.JSON(status, APIResponse{
        StatusCode: status,
        Message: message,
        Data: data,
    })
}
