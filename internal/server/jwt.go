package server

import (
    "errors"
    "net/http"
    "strings"
    "time"

    "github.com/golang-jwt/jwt/v5"
    "github.com/labstack/echo/v4"
)

var jwtSecret = []byte("secret-key")

type JwtCustomClaims struct {
    UserID int `json:"user_id"`
    jwt.RegisteredClaims
}

func GenerateToken(userID int) (string, error) {
    claims := &JwtCustomClaims{
        UserID: userID,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)), // токен живёт 1 час
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        authHeader := c.Request().Header.Get("Authorization")

        if authHeader == "" {
            return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Отсутствует токен авторизации"})
        }

        parts := strings.Split(authHeader, " ")

        if len(parts) != 2 || parts[0] != "Bearer" {
            return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Некорректный заголовок авторизации"})
        }

        token, err := jwt.ParseWithClaims(parts[1], &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, errors.New("unexpected signing method")
            }

            return jwtSecret, nil
        })

        if err != nil || !token.Valid {
            return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Некорректный токен авторизации"})
        }

        claims := token.Claims.(*JwtCustomClaims)
        c.Set("user_id", claims.UserID)

        return next(c)
    }
}
