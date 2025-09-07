package server

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

func RegisterProfileRoutes(e *echo.Echo) {
    e.GET("/profile", func(c echo.Context) error {
        userID := c.Get("user_id").(int)

        return c.JSON(http.StatusOK, map[string]interface{}{
            "id":   userID,
            "name": "Demo User",
        })
    }, JWTMiddleware)
}
