package errors

import "errors"

var (
    ErrIncorrectJson      = errors.New("Некорректный JSON")
    ErrInvalidCredentials = errors.New("Неверные учетные данные")
    ErrTokenGeneration    = errors.New("Не удалось сгенерировать токен")
)
