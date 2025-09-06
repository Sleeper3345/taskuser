package server

import (
    "net/http"
    "strconv"
    "errors"

    "github.com/labstack/echo/v4"
    "taskuser/internal/domain/tasks/models"
    "taskuser/internal/service/taskservice"
    domainErrors "taskuser/internal/domain/tasks/errors"
)

func handleTaskError(c echo.Context, err error) error {
    switch {
    case errors.Is(err, domainErrors.ErrIncorrectId):
        return c.JSON(http.StatusBadRequest, map[string]string{"error": domainErrors.ErrIncorrectId.Error()})
    case errors.Is(err, domainErrors.ErrTaskNotFound):
        return c.JSON(http.StatusNotFound, map[string]string{"error": domainErrors.ErrTaskNotFound.Error()})
    case errors.Is(err, domainErrors.ErrIncorrectJson):
        return c.JSON(http.StatusBadRequest, map[string]string{"error": domainErrors.ErrIncorrectJson.Error()})
    default:
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Внутренняя ошибка сервера"})
    }
}

func RegisterTaskRoutes(e *echo.Echo, service *taskservice.TaskService) {
    e.GET("/tasks", func(c echo.Context) error {
        return c.JSON(http.StatusOK, service.Repository().GetAll())
    })

    e.GET("/tasks/:id", func(c echo.Context) error {
        id, err := strconv.Atoi(c.Param("id"))

        if err != nil {
            return handleTaskError(c, domainErrors.ErrIncorrectId)
        }

        task, err := service.Repository().GetByID(id)

        if err != nil {
            return handleTaskError(c, domainErrors.ErrTaskNotFound)
        }

        return c.JSON(http.StatusOK, task)
    })

    e.POST("/tasks", func(c echo.Context) error {
        var dto models.CreateTask

        if err := c.Bind(&dto); err != nil {
            return handleTaskError(c, domainErrors.ErrIncorrectJson)
        }
    
        if err := c.Validate(&dto); err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
        }
    
        task := models.Task{
            Title:       dto.Title,
            Description: dto.Description,
            Status:      dto.Status,
        }
    
        created := service.Repository().Create(task)

        return c.JSON(http.StatusCreated, created)
    })

    e.PUT("/tasks/:id", func(c echo.Context) error {
        id, err := strconv.Atoi(c.Param("id"))

        if err != nil {
            return handleTaskError(c, domainErrors.ErrIncorrectId)
        }
    
        var dto models.UpdateTask

        if err := c.Bind(&dto); err != nil {
            return handleTaskError(c, domainErrors.ErrIncorrectJson)
        }
    
        if err := c.Validate(&dto); err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
        }
    
        existing, err := service.Repository().GetByID(id)

        if err != nil {
            return handleTaskError(c, domainErrors.ErrTaskNotFound)
        }
    
        if dto.Title != nil {
            existing.Title = *dto.Title
        }

        if dto.Description != nil {
            existing.Description = *dto.Description
        }

        if dto.Status != nil {
            existing.Status = *dto.Status
        }
    
        updated, _ := service.Repository().Update(id, *existing)

        return c.JSON(http.StatusOK, updated)
    })

    e.DELETE("/tasks/:id", func(c echo.Context) error {
        id, err := strconv.Atoi(c.Param("id"))
        
        if err != nil {
            return handleTaskError(c, domainErrors.ErrIncorrectId)
        }

        if err := service.Repository().Delete(id); err != nil {
            return handleTaskError(c, domainErrors.ErrTaskNotFound)
        }

        return c.NoContent(http.StatusNoContent)
    })
}
