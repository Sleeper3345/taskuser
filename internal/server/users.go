package server

import (
    "net/http"
    "strconv"
    "errors"

    "github.com/labstack/echo/v4"
    "taskuser/internal/domain/users/models"
    "taskuser/internal/service/userservice"
    domainErrors "taskuser/internal/domain/users/errors"
)

func handleUserError(c echo.Context, err error) error {
    switch {
    case errors.Is(err, domainErrors.ErrIncorrectId):
        return c.JSON(http.StatusBadRequest, map[string]string{"error": domainErrors.ErrIncorrectId.Error()})
    case errors.Is(err, domainErrors.ErrUserNotFound):
        return c.JSON(http.StatusNotFound, map[string]string{"error": domainErrors.ErrUserNotFound.Error()})
    case errors.Is(err, domainErrors.ErrIncorrectJson):
        return c.JSON(http.StatusBadRequest, map[string]string{"error": domainErrors.ErrIncorrectJson.Error()})
    default:
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Внутренняя ошибка сервера"})
    }
}

func RegisterUserRoutes(e *echo.Echo, service *userservice.UserService) {
    e.GET("/users", func(c echo.Context) error {
        return c.JSON(http.StatusOK, service.Repository().GetAll())
    })

    e.GET("/users/:id", func(c echo.Context) error {
        id, err := strconv.Atoi(c.Param("id"))

        if err != nil {
            return handleUserError(c, domainErrors.ErrIncorrectId)
        }

        user, err := service.Repository().GetByID(id)

        if err != nil {
            return handleUserError(c, domainErrors.ErrUserNotFound)
        }

        return c.JSON(http.StatusOK, user)
    })

    e.POST("/users", func(c echo.Context) error {
        var dto models.CreateUser

        if err := c.Bind(&dto); err != nil {
            return handleUserError(c, domainErrors.ErrIncorrectJson)
        }
    
        if err := c.Validate(&dto); err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
        }
    
        user := models.User{
            Name:     dto.Name,
            Email:    dto.Email,
            Password: dto.Password,
        }
    
        created := service.Repository().Create(user)

        return c.JSON(http.StatusCreated, created)
    })

    e.PUT("/users/:id", func(c echo.Context) error {
        id, err := strconv.Atoi(c.Param("id"))

        if err != nil {
            return handleUserError(c, domainErrors.ErrIncorrectId)
        }
    
        var dto models.UpdateUser

        if err := c.Bind(&dto); err != nil {
            return handleUserError(c, domainErrors.ErrIncorrectJson)
        }
    
        if err := c.Validate(&dto); err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
        }
    
        existing, err := service.Repository().GetByID(id)

        if err != nil {
            return handleUserError(c, domainErrors.ErrUserNotFound)
        }
    
        if dto.Name != nil {
            existing.Name = *dto.Name
        }

        if dto.Email != nil {
            existing.Email = *dto.Email
        }

        if dto.Password != nil {
            existing.Password = *dto.Password
        }
    
        updated, _ := service.Repository().Update(id, *existing)

        return c.JSON(http.StatusOK, updated)
    })

    e.DELETE("/users/:id", func(c echo.Context) error {
        id, err := strconv.Atoi(c.Param("id"))

        if err != nil {
            return handleUserError(c, domainErrors.ErrIncorrectId)
        }

        if err := service.Repository().Delete(id); err != nil {
            return handleUserError(c, domainErrors.ErrUserNotFound)
        }

        return c.NoContent(http.StatusNoContent)
    })
}
