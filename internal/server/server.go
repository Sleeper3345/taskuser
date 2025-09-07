package server

import (
    "github.com/go-playground/validator/v10"
    "github.com/labstack/echo/v4"
    "taskuser/internal/repository/inmemory"
    "taskuser/internal/service/taskservice"
    "taskuser/internal/domain/tasks/models"
    "taskuser/internal/service/userservice"
)

type CustomValidator struct {
    validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
    return cv.validator.Struct(i)
}

func New() *echo.Echo {
    e := echo.New()

    RegisterAuthRoutes(e)
    RegisterProfileRoutes(e)

    v := validator.New()

    v.RegisterValidation("status", func(fl validator.FieldLevel) bool {
        status := models.Status(fl.Field().String())
        return status.IsValid()
    })

    e.Validator = &CustomValidator{validator: v}

    repo := inmemory.NewTaskRepository()
    service := taskservice.NewTaskService(repo)
    RegisterTaskRoutes(e, service)

    userRepo := inmemory.NewUserRepository()
    userService := userservice.NewUserService(userRepo)
    RegisterUserRoutes(e, userService)

    return e
}
