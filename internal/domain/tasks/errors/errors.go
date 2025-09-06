package errors

import "errors"

var (
    ErrTaskNotFound  = errors.New("Задача не найдена")
    ErrIncorrectId   = errors.New("Некорректный ID")
	ErrIncorrectJson = errors.New("Некорректный JSON")
)
