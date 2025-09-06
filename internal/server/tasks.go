package server

import (
    "net/http"
    "strconv"

    "github.com/labstack/echo/v4"
    "taskuser/internal/domain/tasks/models"
    "taskuser/internal/service/taskservice"
)

func RegisterTaskRoutes(e *echo.Echo, service *taskservice.TaskService) {
    e.GET("/tasks", func(c echo.Context) error {
        return c.JSON(http.StatusOK, service.Repository().GetAll())
    })

    e.GET("/tasks/:id", func(c echo.Context) error {
        id, err := strconv.Atoi(c.Param("id"))

        if err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"error": "Некорректный id"})
        }

        task, err := service.Repository().GetByID(id)

        if err != nil {
            return c.JSON(http.StatusNotFound, map[string]string{"error": "Запись не найдена"})
        }

        return c.JSON(http.StatusOK, task)
    })

    e.POST("/tasks", func(c echo.Context) error {
        var dto models.CreateTask

        if err := c.Bind(&dto); err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"error": "Некорректный JSON"})
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
            return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
        }
    
        var dto models.UpdateTask

        if err := c.Bind(&dto); err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"error": "Некорректный JSON"})
        }
    
        if err := c.Validate(&dto); err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
        }
    
        existing, err := service.Repository().GetByID(id)

        if err != nil {
            return c.JSON(http.StatusNotFound, map[string]string{"error": "Запись не найдена"})
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
            return c.JSON(http.StatusBadRequest, map[string]string{"error": "Некорректный id"})
        }

        if err := service.Repository().Delete(id); err != nil {
            return c.JSON(http.StatusNotFound, map[string]string{"error": "Запись не найдена"})
        }

        return c.NoContent(http.StatusNoContent)
    })
}
