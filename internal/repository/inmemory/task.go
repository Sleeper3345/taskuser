package inmemory

import (
    "errors"
    "taskuser/internal/domain/tasks/models"
)

type TaskRepository struct {
    tasks  []models.Task
    nextID int
}

func NewTaskRepository() *TaskRepository {
    return &TaskRepository{
        tasks:  make([]models.Task, 0),
        nextID: 1,
    }
}

func (r *TaskRepository) GetAll() []models.Task {
    return r.tasks
}

func (r *TaskRepository) GetByID(id int) (*models.Task, error) {
    for _, t := range r.tasks {
        if t.ID == id {
            return &t, nil
        }
    }

    return nil, errors.New("Запись не найдена")
}

func (r *TaskRepository) Create(task models.Task) models.Task {
    task.ID = r.nextID
    r.nextID++
    r.tasks = append(r.tasks, task)

    return task
}

func (r *TaskRepository) Update(id int, updated models.Task) (*models.Task, error) {
    for i, t := range r.tasks {
        if t.ID == id {
            updated.ID = id
            r.tasks[i] = updated
            
            return &r.tasks[i], nil
        }
    }

    return nil, errors.New("Запись не найдена")
}

func (r *TaskRepository) Delete(id int) error {
    for i, t := range r.tasks {
        if t.ID == id {
            r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
            return nil
        }
    }

    return errors.New("Запись не найдена")
}
