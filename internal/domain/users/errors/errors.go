package errors

import "errors"

var (
	ErrUserNotFound  = errors.New("Пользователь не найден")
	ErrIncorrectId   = errors.New("Некорректный ID")
	ErrIncorrectJson = errors.New("Некорректный JSON")
)
