package server

import (
    "net/http"
    "errors"

    "github.com/labstack/echo/v4"
    authErrors "taskuser/internal/domain/auth/errors"
)

const demoUsername = "admin"
const demoPassword = "password"

func handleAuthError(c echo.Context, err error) error {
    switch {
    case errors.Is(err, authErrors.ErrIncorrectJson):
        return c.JSON(http.StatusBadRequest, map[string]string{"error": authErrors.ErrIncorrectJson.Error()})
    case errors.Is(err, authErrors.ErrInvalidCredentials):
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": authErrors.ErrInvalidCredentials.Error()})
    case errors.Is(err, authErrors.ErrTokenGeneration):
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": authErrors.ErrTokenGeneration.Error()})
    default:
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Внутренняя ошибка сервера"})
    }
}

func RegisterAuthRoutes(e *echo.Echo) {
    e.POST("/login", func(c echo.Context) error {
        username, password, ok := c.Request().BasicAuth()

        if !ok {
            c.Response().Header().Set("WWW-Authenticate", "Basic realm=\"Restricted\"")
            return handleAuthError(c, authErrors.ErrInvalidCredentials)
        }

        if username != demoUsername || password != demoPassword {
            return handleAuthError(c, authErrors.ErrInvalidCredentials)
        }

        token, err := GenerateToken(1)
		
        if err != nil {
            return handleAuthError(c, authErrors.ErrTokenGeneration)
        }

        return c.JSON(http.StatusOK, map[string]string{"token": token})
    })
}
