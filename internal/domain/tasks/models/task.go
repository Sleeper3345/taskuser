package models

type Status string

const (
    StatusNew       Status = "Новая"
    StatusInProcess Status = "В процессе"
    StatusDone      Status = "Завершена"
)

type Task struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Status      Status `json:"status"`
}

type CreateTask struct {
    Title       string `json:"title" validate:"required"`
    Description string `json:"description"`
    Status      Status `json:"status" validate:"status"`
}

type UpdateTask struct {
    Title       *string `json:"title,omitempty" validate:"omitempty"`
    Description *string `json:"description,omitempty"`
    Status      *Status `json:"status,omitempty" validate:"omitempty,status"`
}

func (s Status) IsValid() bool {
    switch s {
    case StatusNew, StatusInProcess, StatusDone:
        return true
    }
    
    return false
}
